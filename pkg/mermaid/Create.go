package mermaid

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/compose-spec/compose-go/types"
)

type Volume struct {
	VolumeConfig types.ServiceVolumeConfig
	ServiceName  string
}

func Create(project *types.Project, templatePath string) (bytes.Buffer, error) {
	return process(sort(project), templatePath)
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

	return volumes
}

func process(volumes map[string][]Volume, templatePath string) (bytes.Buffer, error) {
	var buf bytes.Buffer

	data, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return buf, err
	}

	tn, err := template.New("template").Parse(string(data))
	if err != nil {
		return buf, err
	}

	err = tn.Execute(&buf, volumes)
	return buf, err
}
