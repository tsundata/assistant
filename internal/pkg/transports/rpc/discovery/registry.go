package discovery

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type MultiServiceDiscovery struct {
	registry   string
	timeout    time.Duration
	lastUpdate time.Time
	service    string
	kv         map[string]string
	pairs      []*client.KVPair
	chans      []chan []*client.KVPair
	mu         sync.Mutex
	// -1 means it always retry to watch until zookeeper is ok, 0 means no retry.
	RetriesAfterWatchFailed int

	filter client.ServiceDiscoveryFilter

	stopCh chan struct{}
}

func NewMultiServiceDiscovery(service string, registry string) client.ServiceDiscovery {
	d := &MultiServiceDiscovery{service: service, registry: registry}
	d.stopCh = make(chan struct{})
	d.filter = func(kvp *client.KVPair) bool {
		return strings.ToLower(kvp.Value) == strings.ToLower(service)
	}

	ps, err := d.refresh()
	if err != nil {
		log.Printf("cannot get services of from registry: %v, err: %v", service, err)
	}

	pairs := make([]*client.KVPair, 0, len(ps))
	for _, p := range ps {
		pair := &client.KVPair{Key: p.Key, Value: p.Value}
		if d.filter != nil && !d.filter(pair) {
			continue
		}
		pairs = append(pairs, pair)
	}
	d.pairs = pairs
	d.RetriesAfterWatchFailed = -1
	go d.watch()

	return d
}

var _ client.ServiceDiscovery = (*MultiServiceDiscovery)(nil)

func (m *MultiServiceDiscovery) GetServices() []*client.KVPair {
	return m.pairs
}

func (m *MultiServiceDiscovery) WatchService() chan []*client.KVPair {
	m.mu.Lock()
	defer m.mu.Unlock()

	ch := make(chan []*client.KVPair, 10)
	m.chans = append(m.chans, ch)
	return ch
}

func (m *MultiServiceDiscovery) RemoveWatcher(ch chan []*client.KVPair) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var chans []chan []*client.KVPair
	for _, c := range m.chans {
		if c == ch {
			continue
		}

		chans = append(chans, c)
	}

	m.chans = chans
}

func (m *MultiServiceDiscovery) Clone(servicePath string) client.ServiceDiscovery {
	return NewMultiServiceDiscovery(servicePath, m.registry)
}

func (m *MultiServiceDiscovery) SetFilter(filter client.ServiceDiscoveryFilter) {
	m.filter = filter
}

func (m *MultiServiceDiscovery) Close() {
	close(m.stopCh)
}

func (m *MultiServiceDiscovery) watch() {
	for {
		var err error
		var c <-chan []*client.KVPair
		var tempDelay time.Duration

		retry := m.RetriesAfterWatchFailed
		for m.RetriesAfterWatchFailed < 0 || retry >= 0 {
			c, err = m.watchRefresh(nil)
			if err != nil {
				if m.RetriesAfterWatchFailed > 0 {
					retry--
				}
				if tempDelay == 0 {
					tempDelay = 1 * time.Second
				} else {
					tempDelay *= 2
				}
				if max := 30 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("can not watchtree (with retry %d, sleep %v): %s: %v", retry, tempDelay, m.service, err)
				time.Sleep(tempDelay)
				continue
			}
			break
		}

		if err != nil {
			log.Printf("can't watch %s: %v", m.service, err)
			return
		}

	readChanges:
		for {
			select {
			case <-m.stopCh:
				log.Println("discovery has been closed")
				return
			case ps := <-c:
				if ps == nil {
					break readChanges
				}
				var pairs []*client.KVPair
				for _, p := range ps {
					pair := &client.KVPair{Key: p.Key, Value: p.Value}
					if m.filter != nil && !m.filter(pair) {
						continue
					}
					pairs = append(pairs, pair)
				}
				m.pairs = pairs

				m.mu.Lock()
				for _, ch := range m.chans {
					ch := ch
					go func() {
						defer func() {
							recover()
						}()
						select {
						case ch <- pairs:
						case <-time.After(time.Minute):
							log.Println("chan is full and new change has been dropped")
						}
					}()
				}
				m.mu.Unlock()
			}

			time.Sleep(10 * time.Second)
		}

		log.Println("chan is closed and will rewatch")

		time.Sleep(15 * time.Second)
	}
}

func (m *MultiServiceDiscovery) refresh() ([]*client.KVPair, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	log.Println("rpc registry: refresh servers from registry", m.registry)

	c, err := net.Dial("tcp", m.registry)
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return []*client.KVPair{}, err
	}

	err = c.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return []*client.KVPair{}, err
	}

	_, err = fmt.Fprintf(c, "GET\n")
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return []*client.KVPair{}, err
	}

	resp, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return []*client.KVPair{}, err
	}

	resp = strings.Trim(resp, "\n")
	servers := strings.Split(resp, ",")
	pairs := make([]*client.KVPair, 0, len(servers))
	for _, server := range servers {
		if strings.TrimSpace(server) != "" {
			item := strings.Split(server, "|")
			if len(item) != 2 {
				return []*client.KVPair{}, errors.New("ServerItem error")
			}
			pair := &client.KVPair{Key: item[1], Value: item[0]}
			if m.filter != nil && !m.filter(pair) {
				continue
			}
			pairs = append(pairs, pair)
		}
	}
	m.lastUpdate = time.Now()

	return pairs, nil
}

func (m *MultiServiceDiscovery) watchRefresh(stopCh <-chan struct{}) (<-chan []*client.KVPair, error) {
	watchCh := make(chan []*client.KVPair)

	go func() {
		defer close(watchCh)

		for {
			select {
			case <-stopCh:
				return
			default:
			}

			pairs, err := m.refresh()
			if err != nil {
				return
			}

			// Return children KV pairs to the channel
			var kvpairs []*client.KVPair
			for _, pair := range pairs {
				if m.filter != nil && !m.filter(pair) {
					continue
				}
				kvpairs = append(kvpairs, &client.KVPair{
					Key:   pair.Key,
					Value: pair.Value,
				})
			}

			watchCh <- kvpairs
		}
	}()

	return watchCh, nil
}
