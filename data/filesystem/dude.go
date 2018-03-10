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

	for _, file := range files {
		f, err := os.Open(file.Name())
		if err != nil {
			return []riley.Dude{}, err
		}
	}

}

//func GetDude(location string)(riley.Dude, error)
