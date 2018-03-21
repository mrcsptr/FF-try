package filesystem

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/moxar/riley"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// AllDudes gets the dudes and put them in the appropriate structure.
func AllDudes(location string) ([]riley.Dude, error) {
	// ... load dudes.
	files, err := ioutil.ReadDir(location)
	if err != nil {
		return nil, err
	}
	var dudes []riley.Dude
	// parsing all files in directory
	for _, file := range files {
		d, err := GetDude(file.Name())
		if err != nil {
			return []riley.Dude{}, err
		}
		dudes = append(dudes, d)
	}
	return dudes, err
}

// GetDude parses the content of a dude file, and returns an error if smething went wrong along the way.
func GetDude(location string) (riley.Dude, error) {
	file, err := os.Open(location)
	if err != nil {
		return riley.Dude{}, err
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	var results []riley.DudeResult
	for scanner.Scan() {
		r, err := parseResult(scanner.Text())
		switch err {
		case EmptyLineErr:
			continue
		default:
			return riley.Dude{}, err
		}
		results = append(results, r)
	}
	if err := scanner.Err(); err != nil {
		return riley.Dude{}, err
	}
	var d riley.Dude
	d.Name = filepath.Clean(location)
	d.Results = results
	return d, err
}
// EmptyLineErr allows to make a difference between error due to empty line or comments and errors due to bad fillings.
var EmptyLineErr = errors.New("EmptyLineErr")

// parseResult exploits the data contained in the dude files, and returns an error if a dude has been badly filled.
func parseResult(content string) (riley.DudeResult, error) {

	entry := strings.Fields(content)
	if len(entry) == 0 {
		return riley.DudeResult{}, EmptyLineErr
	}
	if entry[0][0] == '#' {
		return riley.DudeResult{}, EmptyLineErr
	}

	p, err := parsePosition(entry[0])
	if err != nil {
		return riley.DudeResult{}, err
	}

	score, err := riley.ParseScore(entry[1])
	if err != nil {
		return riley.DudeResult{}, err
	}

	var r riley.DudeResult
	r.Score = score
	r.Position = p
	return r, nil
}

// parsePosition checks the position occupied by the dude for the selected entry, and returns an error if character is missing or not what was expected
func parsePosition(content string) (string, error) {
	switch {
	case strings.EqualFold(content, "a"):
		return "A", nil
	case strings.EqualFold(content, "m"):
		return "M", nil
	case strings.EqualFold(content, "d"):
		return "D", nil
	default:
		return "", fmt.Errorf("invalid position '%s'", content)
	}
}
