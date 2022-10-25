package parser

import (
	"log"

	"github.com/gachampiat/compose-diag/pkg/mermaid"
	"github.com/gachampiat/compose-diag/pkg/nwdiag"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

func getSection(config map[string]interface{}, key string) map[string]interface{} {
	section, ok := config[key]
	if !ok {
		return make(map[string]interface{})
	}
	return section.(map[string]interface{})
}

func Parse(content []byte) {
	project, err := loader.Load(types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{
			{Filename: "filename.yml", Content: content},
		},
	})
	if err != nil {
		log.Panic(err)
	}

	nwdiag.Create(project)
	mermaid.Create(project)
}
