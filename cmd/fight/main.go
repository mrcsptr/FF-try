package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Define configuration filepath.
	var configPath string
	flag.StringVar(&configPath, "c", "config.json", "configuration filepath")

	flag.Parse()

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("unable to load configuration:", err)
	}

	fmt.Println(string(b))
}
