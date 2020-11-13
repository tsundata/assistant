package registry

import (
	"log"
	"net/http"
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
	defaultPath    = "/_rpc_/registry"
	defaultTimeout = time.Minute * 5
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

func (r *Registry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Set("X-RPC-Servers", strings.Join(r.aliveServers(), ","))
	case "POST":
		servicePath := req.Header.Get("X-RPC-Path")
		if servicePath == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		addr := req.Header.Get("X-RPC-Addr")
		if addr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.putServer(servicePath, addr)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (r *Registry) HandleHTTP(registryPath string) {
	http.Handle(registryPath, r)
	log.Println("rpc registry path:", registryPath)
}

func HandleHTTP() {
	DefaultRegister.HandleHTTP(defaultPath)
}

func Heartbeat(registry, servicePath, addr string, duration time.Duration) {
	if duration == 0 {
		duration = defaultTimeout - time.Duration(1)*time.Minute
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
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", registry, nil)
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return err
	}
	req.Header.Set("X-RPC-Path", servicePath)
	req.Header.Set("X-RPC-Addr", addr)
	if _, err := httpClient.Do(req); err != nil {
		log.Println("rpc server: heart beat err:", err)
		return err
	}
	return nil
}