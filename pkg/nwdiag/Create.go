package nwdiag

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/compose-spec/compose-go/types"
)

type Subnet struct {
	Network  types.NetworkConfig
	Services []types.ServiceConfig
}

type Subnets struct {
	Configuration map[string]Subnet
	Groups        map[string][]string
}

func Create(project *types.Project) error {

	process(sort(project))
	return nil
}

func process(subnets Subnets) {
	data, err := ioutil.ReadFile("assets/nwdiag.tpl")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	tn, err := template.New("template").Parse(string(data))
	if err != nil {
		log.Panic(err)
	}

	err = tn.Execute(os.Stdout, subnets)
	if err != nil {
		log.Panic(err)
	}
}

func sort(project *types.Project) Subnets {
	var subnets = Subnets{map[string]Subnet{}, map[string][]string{}}

	// Initialisation des subnets
	for _, network := range project.Networks {
		// remove _
		subnets.Configuration[network.Name[1:]] = Subnet{
			network,
			[]types.ServiceConfig{},
		}
	}

	// Go through all services and add service into
	// the correct network
	for _, service := range project.AllServices() {

		// If the service is configured to be
		// in network mode
		if service.NetworkMode != "" {
			mode := strings.Split(service.NetworkMode, ":")
			if mode[0] == "host" {
				log.Println("Host network not yet supported")
			} else {
				if entry, ok := subnets.Groups[mode[1]]; ok {

					// Then we modify the copy
					entry = append(subnets.Groups[mode[1]], service.Name)

					// Then we reassign map entry
					subnets.Groups[mode[1]] = entry
				} else {
					subnets.Groups[mode[1]] = []string{service.Name}
				}
			}
		}

		for name := range service.Networks {
			if entry, ok := subnets.Configuration[name]; ok {
				entry.Services = append(subnets.Configuration[name].Services, service)
				subnets.Configuration[name] = entry
			}

		}
	}
	return subnets
}

func validate(subnets Subnets) error {
	var err error

	for _, subnet := range subnets.Configuration {
		if len(subnet.Network.Ipam.Config) > 1 {
			err = fmt.Errorf("Docker-compose file not supported. Network %s contains more than 1 IPAM Config", subnet.Network.Name)
			break
		}
	}

	return err

}
