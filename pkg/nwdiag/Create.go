package nwdiag

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/compose-spec/compose-go/types"
)

type Subnet struct {
	Network  types.NetworkConfig
	Services []types.ServiceConfig
}

func Create(project *types.Project) error {

	process(sort(project))
	return nil
}

func process(subnets []Subnet) {
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

func sort(project *types.Project) []Subnet {
	var subnets = []Subnet{}

	for _, network := range project.Networks {
		log.Println("Processing Network : ", network.Name)
		var services = []types.ServiceConfig{}
		for _, service := range project.AllServices() {
			if _, ok := service.Networks[network.Name[1:]]; ok {
				services = append(services, service)
			}
		}
		subnets = append(subnets, Subnet{network, services})
	}

	return subnets
}

func validate(subnets []Subnet) error {
	var err error

	for _, subnet := range subnets {
		if len(subnet.Network.Ipam.Config) > 1 {
			err = fmt.Errorf("Docker-compose file not supported. Network %s contains more than 1 IPAM Config", subnet.Network.Name)
			break
		}
	}

	return err

}
