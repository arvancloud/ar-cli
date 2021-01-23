package main

import (
	"github.com/ebrahimahmadi/ar-cli/cmd"
	"github.com/ebrahimahmadi/ar-cli/config"
	"log"
)

func main() {
	_, err := config.LoadConfigFile()

	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
