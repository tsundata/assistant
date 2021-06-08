package discovery

import (
	"fmt"
	"github.com/hashicorp/consul/api"
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
	consulAddress := os.Getenv("CONSUL_ADDRESS") // todo
	client, err := api.NewClient(&api.Config{
		Address: consulAddress,
		Scheme:  "http", // todo
	})
	if err != nil {
		fmt.Printf("error create consul client: %v\n", err)
		return
	}

	agent := client.Agent()
	interval := time.Duration(10) * time.Second
	deregister := time.Duration(1) * time.Minute

	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v/%v:%v", cs.Name, cs.IP, cs.Port), // name of service node
		Name:    cs.Name,                                          // Service Name
		Tags:    cs.Tag,                                           // tag, can be empty
		Port:    cs.Port,                                          // Service Port
		Address: cs.IP,                                            // Service IP
		Check: &api.AgentServiceCheck{ // Health Examination
			Interval:                       interval.String(),                                // health check interval
			GRPC:                           fmt.Sprintf("%v:%v/%v", cs.IP, cs.Port, cs.Name), // grpc support, address to perform health check, service will be passed to Health. Check function
			DeregisterCriticalServiceAfter: deregister.String(),                              // logout time, equivalent to expiration time
		},
	}

	fmt.Printf("registing to %v\n", ca)
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Printf("Service Register error\n%v", err)
		return
	}
}
