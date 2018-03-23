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
		if err == EmptyLineErr {
			continue
        } else {
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
// EmptyLineErr is returned when the line to parse is empty or only contains a comment
var EmptyLineErr = errors.New("EmptyLineErr")

// parseResult exploits the data contained in the dude files, and returns an error if a dude has been badly filled.
func parseResult(content string) (riley.DudeResult, error) {

	entries := strings.Fields(content)
	switch len(entries) {
        case 0:
            return riley.DudeResult{}, EmptyLineErr
        case 2:
            continue
        default:
            return riley.DudeResult{}, fmt.Errorf("invalid entries")
	}
	if entries[0][0] == '#' {
		return riley.DudeResult{}, EmptyLineErr
	}

	p, err := parsePosition(entries[0])
	if err != nil {
		return riley.DudeResult{}, err
	}

	score, err := riley.ParseScore(entries[1])
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
