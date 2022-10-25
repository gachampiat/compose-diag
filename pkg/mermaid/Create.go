package mermaid

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/compose-spec/compose-go/types"
)

type Volume struct {
	VolumeConfig types.ServiceVolumeConfig
	ServiceName  string
}

func Create(project *types.Project) error {
	var err error

	process(sort(project))

	return err
}

func sort(project *types.Project) map[string][]Volume {
	var volumes = map[string][]Volume{}

	for _, service := range project.AllServices() {
		for _, volume := range service.Volumes {
			if entry, ok := volumes[volume.Source]; ok {

				// Then we modify the copy
				entry = append(volumes[volume.Source], Volume{volume, service.Name})

				// Then we reassign map entry
				volumes[volume.Source] = entry
			} else {
				volumes[volume.Source] = []Volume{{volume, service.Name}}
			}
		}
	}

	log.Println(volumes)
	return volumes
}

func process(volumes map[string][]Volume) {
	data, err := ioutil.ReadFile("assets/mermaid.tpl")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	tn, err := template.New("template").Parse(string(data))
	if err != nil {
		log.Panic(err)
	}

	err = tn.Execute(os.Stdout, volumes)
	if err != nil {
		log.Panic(err)
	}
}
