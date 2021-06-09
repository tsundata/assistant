package discovery

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"os"
	"regexp"
	"sync"
)

const defaultPort = "7001"

var (
	errMissingAddr = errors.New("consul resolver: missing address")

	errAddrMisMatch = errors.New("consul resolver: invalied uri")

	// errEndsWithColon = errors.New("consul resolver: missing port after port-separator colon")

	regexConsul, _ = regexp.Compile("^([A-z0-9.]+)(:[0-9]{1,5})?/([A-z_]+)$")
)

func RegisterBuilder() {
	fmt.Println("calling consul init")
	resolver.Register(NewBuilder())
}

type consulBuilder struct{}

type consulResolver struct {
	address              string
	wg                   sync.WaitGroup
	cc                   resolver.ClientConn
	name                 string
	disableServiceConfig bool
	lastIndex            uint64
}

func NewBuilder() resolver.Builder {
	return &consulBuilder{}
}

func (cb *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	fmt.Printf("calling consul build\n")
	fmt.Printf("target: %v\n", target)
	host, port, name, err := parseTarget(fmt.Sprintf("%s/%s", target.Authority, target.Endpoint))
	if err != nil {
		return nil, err
	}

	cr := &consulResolver{
		address:              fmt.Sprintf("%s%s", host, port),
		cc:                   cc,
		name:                 name,
		disableServiceConfig: opts.DisableServiceConfig,
		lastIndex:            0,
	}

	cr.wg.Add(1)
	go cr.watcher()
	return cr, nil
}

func (cb *consulBuilder) Scheme() string {
	return "consul"
}

func (cr *consulResolver) watcher() {
	fmt.Println("calling consul watcher")
	consulAddress := os.Getenv("CONSUL_ADDRESS") // todo
	client, err := api.NewClient(&api.Config{
		Address: consulAddress,
		Scheme:  "http", // todo
	})
	if err != nil {
		fmt.Printf("error create consul client: %v\n", err)
		return
	}

	for {
		services, metaInfo, err := client.Health().Service(cr.name, cr.name, true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			fmt.Printf("error retrieving instances from Consul: %v\n", err)
		}

		cr.lastIndex = metaInfo.LastIndex
		var newAddr []resolver.Address
		for _, service := range services {
			addr := fmt.Sprintf("%v:%v", service.Service.Address, service.Service.Port)
			newAddr = append(newAddr, resolver.Address{Addr: addr})
		}
		fmt.Println("adding service addr")
		fmt.Printf("newAddrs: %v\n", newAddr)
		cr.cc.NewAddress(newAddr)       // nolint
		cr.cc.NewServiceConfig(cr.name) // nolint
	}
}

func (cr *consulResolver) ResolveNow(_ resolver.ResolveNowOptions) {
}

func (cr *consulResolver) Close() {
}

func parseTarget(target string) (host, port, name string, err error) {
	fmt.Printf("target uri: %v\n", target)
	if target == "" {
		return "", "", "", errMissingAddr
	}
	if !regexConsul.MatchString(target) {
		return "", "", "", errAddrMisMatch
	}

	groups := regexConsul.FindStringSubmatch(target)
	host = groups[1]
	port = groups[2]
	name = groups[3]
	if port == "" {
		port = defaultPort
	}
	return
}
