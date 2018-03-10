package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/moxar/riley"
)

func AllDudes(location string) ([]riley.Dude, error) {
	// ... load dudes.
	files, err := ioutil.ReadDir(location)
	if err != nil {
		return nil , err
	}
var d []Dude
	for _, file := range files {
        g :=GetDude(file.Name())
        d = append(d,g)
	}
}

func GetDude(location string)(riley.Dude, error){
    var d Dude
    d.Name = location
    results, err := os.Open(location)
    if err!= nil {
        return  results, err
    }
    return d, nil
}
