package main

import (
	"flag"
	"fmt"
	"log"
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

	fmt.Println(config)
}
