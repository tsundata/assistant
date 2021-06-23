package discovery

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"os"
	"time"
)

type ConsulService struct {
	IP   string
	Port int
	Tag  []string
	Name string
}

func RegisterService(ca string, cs *ConsulService) {
	//register consul
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

	agent := client.Agent()
	interval := time.Duration(10) * time.Second
	deregister := time.Duration(10) * time.Minute

	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v/%v:%v", cs.Name, cs.IP, cs.Port), // name of service node
		Name:    cs.Name,                                          // Service Name
		Tags:    cs.Tag,                                           // tag, can be empty
		Port:    cs.Port,                                          // Service Port
		Address: cs.IP,                                            // Service IP
		Check: &api.AgentServiceCheck{ // Health Examination
			Interval:                       interval.String(), // health check interval
			TCP:                            fmt.Sprintf("%s:%d", cs.IP, cs.Port),
			DeregisterCriticalServiceAfter: deregister.String(), // logout time, equivalent to expiration time
			//	GRPC: fmt.Sprintf("%v:%v/%v", cs.IP, cs.Port, cs.Name), // grpc support, address to perform health check, service will be passed to Health. Check function
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		log.Printf("Service Register error\n%v", err)
		return
	}
}
