package main

import (
	"io/ioutil"
	"log"

	"github.com/gachampiat/compose-diag/pkg/parser"
)

func main() {
	message, err := ioutil.ReadFile("example/network-docker-compose.yaml")
	if err != nil {
		log.Fatal(err)
	}
	parser.Parse(message)
}
