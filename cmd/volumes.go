package cmd

import (
	"log"
	"path/filepath"

	"github.com/gachampiat/compose-diag/pkg/mermaid"
	"github.com/gachampiat/compose-diag/pkg/parser"

	"github.com/spf13/cobra"
)

var volumeCmd = &cobra.Command{
	Use:     "volumes",
	Aliases: []string{"vol", "v"},
	Short:   "Create a volumes diagram",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sourceAbs, err := filepath.Abs(source)
		if err != nil {
			log.Fatalf("An error appears : %s\n", err)
		}
		templateFileAbs, err := filepath.Abs(templateFile)
		if err != nil {
			log.Fatalf("An error appears : %s\n", err)
		}

		log.Printf("Parsing docker-compose %s\n", sourceAbs)
		project, err := parser.Parse(sourceAbs)
		if err != nil {
			log.Fatalf("An error appears during the docker-compose (%s) parsing.\n%s\n", sourceAbs, err)
		}

		log.Println("Creating mermaid diagram.....")
		buf, err := mermaid.Create(project, templateFileAbs)
		if err != nil {
			log.Fatalf("An error appears during the conversion of docker-compose (%s) into mermaid diagram.\n%s\n", sourceAbs, err)

		}

		log.Println(buf.String())
	},
}

func init() {
	volumeCmd.Flags().StringVarP(&source, "file", "f", "", "Docker-compose file to read from")
	volumeCmd.Flags().StringVarP(&templateFile, "templateFile", "t", "assets/mermaid.tpl", "Go template for create output file")
	volumeCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(volumeCmd)
}
