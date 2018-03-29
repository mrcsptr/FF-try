package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/moxar/riley/data/filesystem"
)

func main() {
	// Define configuration filepath.
	var configPath string
	flag.StringVar(&configPath, "c", "config.json", "configuration filepath")

	flag.Parse()
	// loading configuration
	config, err := NewConfig(configPath)
	if err != nil {
		log.Fatalln("unable to load configuration:", err)
	}

	dudes, err := filesystem.AllDudes(config.DudesLocation)
	if err != nil {
		log.Fatalln("unable to parse dudes:", err)
	}

	teams, err := filesystem.AllTeams(config.TeamsLocation)
	if err != nil {
		log.Fatalln("unable to parse teams:", err)
	}

	fmt.Println(teams)
    fmt.Println(dudes)
}
