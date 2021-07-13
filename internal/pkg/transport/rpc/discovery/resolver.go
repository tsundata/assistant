package discovery

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"log"
	"os"
	"regexp"
	"sync"
	"time"
)

const defaultPort = "7001"

var (
	errMissingAddr = errors.New("consul resolver: missing address")

	errAddrMisMatch = errors.New("consul resolver: invalied uri")

	// errEndsWithColon = errors.New("consul resolver: missing port after port-separator colon")

	regexConsul, _ = regexp.Compile("^([A-z0-9.]+)(:[0-9]{1,5})?/([A-z_]+)$")
)

func RegisterBuilder() {
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
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	consulScheme := os.Getenv("CONSUL_SCHEME")
	consulUsername := os.Getenv("CONSUL_USERNAME")
	consulPassword := os.Getenv("CONSUL_PASSWORD")
	consulToken := os.Getenv("CONSUL_TOKEN")
	client, err := api.NewClient(&api.Config{
		Address: consulAddress,
		Scheme:  consulScheme,
		HttpAuth: &api.HttpBasicAuth{
			Username: consulUsername,
			Password: consulPassword,
		},
		Token: consulToken,
	})
	if err != nil {
		log.Printf("error create consul client: %v\n", err)
		return
	}

	for {
		services, metaInfo, err := client.Health().Service(cr.name, "grpc", true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			log.Printf("error retrieving instances from Consul: %v\n", err)
		}

		cr.lastIndex = metaInfo.LastIndex
		var newAddr []resolver.Address
		for _, service := range services {
			addr := fmt.Sprintf("%v:%v", service.Service.Address, service.Service.Port)
			newAddr = append(newAddr, resolver.Address{Addr: addr})
		}
		cr.cc.NewAddress(newAddr)       // nolint
		cr.cc.NewServiceConfig(cr.name) // nolint
		if len(newAddr) > 0 {
			log.Printf("%s newAddrs: %v\n", cr.name, newAddr)
			time.Sleep(time.Second)
		}
	}
}

func (cr *consulResolver) ResolveNow(_ resolver.ResolveNowOptions) {}

func (cr *consulResolver) Close() {}

func parseTarget(target string) (host, port, name string, err error) {
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
