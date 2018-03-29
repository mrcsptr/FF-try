package filesystem

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/moxar/riley"
)

// AllTeams gets the teams and put them in the appropriate structure.
func AllTeams(location string) ([]riley.Team, error) {
	// ... load teams.
	files, err := ioutil.ReadDir(location)
	if err != nil {
		return nil, err
	}
	var teams []riley.Team
	// parsing all files in directory
	for _, file := range files {
		t, err := GetTeam(location + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	return teams, err
}

// GetTeam parses the content of a team file, and returns an error if something went wrong along the way.
func GetTeam(path string) (riley.Team, error) {
	file, err := os.Open(path)
	if err != nil {
		return riley.Team{}, err
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	var comp []riley.Teammate
	for scanner.Scan() {
		c, err := parseComp(scanner.Text())
		if err != nil {
			if err != ErrEmptyLine {
				return riley.Team{}, err
			}
			continue
		}
		comp = append(comp, c)
	}
	if err := scanner.Err(); err != nil {
		return riley.Team{}, err
	}
	var t riley.Team
	t.Name = filepath.Clean(strings.TrimSuffix(path, filepath.Ext(path)))
	t.Comp = comp
	return t, err
}

// parseResult exploits the data contained in the dude files, and returns an error if a dude has been badly filled.
func parseComp(content string) (riley.Teammate, error) {

	chunks := strings.Fields(content)
	if len(chunks) < 2 {
		if len(chunks) == 0 {
			return riley.Teammate{}, ErrEmptyLine
		}
		return riley.Teammate{}, newInvalidLineError(content)
	}
	if chunks[0][0] == '#' {
		return riley.Teammate{}, ErrEmptyLine
	}



	dude, err := riley.ParseDude(chunks[0])
	if err != nil {
		return riley.Teammate{}, err
	}
	
		p, err := parsePosition(chunks[1])
	if err != nil {
		return riley.Teammate{}, err
	}

	var t riley.Teammate
	t.Dude = dude
	t.Position = p
	return t, nil
}

// parseDude checks the dude in the and search for the data  // about him in AllDudes and returns an error if character is // missing or not what was expected.
func parseDude(content string) (riley.Dude, error) {
	}
}




// parsePosition checks the position occupied by the dude for the selected entry,
// and returns an error if character is missing or not what was expected.
func parsePosition(content string) (string, error) {
	switch {
	case strings.EqualFold(content, "a"):
		return "A", nil
	case strings.EqualFold(content, "m"):
		return "M", nil
	case strings.EqualFold(content, "d"):
		return "D", nil
	default:
		return "", newInvalidPositionError(content)
	}
}
