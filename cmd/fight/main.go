package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/moxar/riley"
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

	teams, err := filesystem.AllTeams(config.TeamsLocation)
	if err != nil {
		log.Fatalln("unable to parse teams:", err)
	}
	
	for _, t := range teams {
        grades := riley.NewGrades()
        for _, m := range t.Teammates{
            dude, err := filesystem.GetDude(config.DudesLocation + "/" + m.Name + ".txt")
            if err != nil {
                log.Fatalln("unable to parse dude:", err)
            }
            
            grades.AddScore(m.Position, dude) 
        }
        fmt.Println("team:", t.Name)
        fmt.Println("players:", t.Teammates)
        fmt.Println("grade:", grades)
        fmt.Println("")
    }
}
