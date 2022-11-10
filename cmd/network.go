package cmd

import (
	"log"
	"path/filepath"

	"github.com/gachampiat/compose-diag/pkg/nwdiag"
	"github.com/gachampiat/compose-diag/pkg/parser"

	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net", "n"},
	Short:   "Create a network diagram",
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

		log.Println("Creating nwdiag.....")
		buf, err := nwdiag.Create(project, templateFileAbs)
		if err != nil {
			log.Fatalf("An error appears during the conversion of docker-compose (%s) into nwdiag diagram.\n%s\n", sourceAbs, err)

		}

		log.Println(buf.String())
	},
}

func init() {
	networkCmd.Flags().StringVarP(&source, "file", "f", "", "Docker-compose file to read from")
	networkCmd.Flags().StringVarP(&templateFile, "templateFile", "t", "assets/nwdiag.tpl", "Go template for create output file")
	networkCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(networkCmd)
}
