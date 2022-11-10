package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compose-diag",
	Short: "compose-diag - create a set of diagrams (network, volume) based on docker-compose file",
	Long: `
 _________                                                         .___.__                
 \_   ___ \  ____   _____ ______   ____  ______ ____             __| _/|__|____     ____  
 /    \  \/ /  _ \ /     \\____ \ /  _ \/  ___// __ \   ______  / __ | |  \__  \   / ___\ 
 \     \___(  <_> )  Y Y  \  |_> >  <_> )___ \\  ___/  /_____/ / /_/ | |  |/ __ \_/ /_/  >
  \______  /\____/|__|_|  /   __/ \____/____  >\___  >         \____ | |__(____  /\___  / 
		 \/             \/|__|              \/     \/               \/         \//_____/

compose-diag facilites documentation generation for a docker-compose file. 
It can:
	* Creates a network diagram to describe the inter-relationship between the containers
	* Volume`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var source string
var templateFile string

func Execute() error {
	return rootCmd.Execute()
}
