package registry

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

type Registry struct {
	timeout time.Duration
	mu      sync.Mutex
	servers map[string]*ServerItem
}

type ServerItem struct {
	ServicePath string
	Addr        string
	start       time.Time
}

const (
	defaultTimeout = 30 * time.Second
)

func New(timeout time.Duration) *Registry {
	return &Registry{
		servers: make(map[string]*ServerItem),
		timeout: timeout,
	}
}

var DefaultRegister = New(defaultTimeout)

func (r *Registry) putServer(servicePath, addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	servicePath = strings.ToLower(servicePath)
	s := r.servers[addr]
	if s == nil {
		r.servers[addr] = &ServerItem{ServicePath: servicePath, Addr: addr, start: time.Now()}
	} else {
		s.start = time.Now()
	}
}

func (r *Registry) aliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	for addr, s := range r.servers {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, s.ServicePath+"|"+addr)
		} else {
			delete(r.servers, addr)
		}
	}
	sort.Strings(alive)
	return alive
}

func (r *Registry) HandleConnection(c net.Conn) {
	defer c.Close()
	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return
	}

	netData = strings.Trim(netData, "\n")
	params := strings.Split(netData, " ")

	if len(params) < 1 {
		_, _ = c.Write([]byte("ERR\n"))
		return
	}

	switch params[0] {
	case "GET":
		_, _ = c.Write([]byte(strings.Join(r.aliveServers(), ",") + "\n"))
		return
	case "POST":
		if len(params) != 3 {
			_, _ = c.Write([]byte("ERR\n"))
			return
		}
		servicePath := params[1]
		if servicePath == "" {
			_, _ = c.Write([]byte("ERR\n"))
			return
		}
		addr := params[2]
		if addr == "" {
			_, _ = c.Write([]byte("ERR\n"))
			return
		}
		r.putServer(servicePath, addr)
		_, _ = c.Write([]byte("OK\n"))
	default:
		_, _ = c.Write([]byte("ERR\n"))
	}
}

func Heartbeat(registry, servicePath, addr string, duration time.Duration) {
	if duration == 0 {
		duration = defaultTimeout - time.Duration(15)*time.Second
	}
	var err error
	err = sendHeartbeat(registry, servicePath, addr)
	go func() {
		t := time.NewTicker(duration)
		for err == nil {
			<-t.C
			err = sendHeartbeat(registry, servicePath, addr)
		}
	}()
}

func sendHeartbeat(registry, servicePath, addr string) error {
	log.Println(servicePath, addr, "send heart beat to registry", registry)

	c, err := net.Dial("tcp", registry)
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return err
	}

	_, err = c.Write([]byte(fmt.Sprintf("POST %s %s\n", servicePath, addr)))
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return err
	}

	return nil
}
