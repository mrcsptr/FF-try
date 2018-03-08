package main

import (
	"io"
)

func AllDudes(location string) ([]riley.Dude, error) {
	// ... load dudes.
	raw, err := io.Reader(location)
	if err != nil {
		return []riley.Dude{}, err
	}
}

// Dude a type holding its characteristics.
type Dude struct {
	Name    string
	Results []DudeResult
}

// DudeResult represents one result, with the position occupied then by the player, and the score obtained.
type DudeResult struct {
	Position string
	Score    int
}

// Definition of the values for a win (V), tie (T), and a loss (L) (unused for now).
const (
	W = 3
	T = 2
	L = 1
)

func (d Dude) ScoreA() int {
	A := d.Score("A")
	return A
}

func (d Dude) ScoreM() int {
	M := d.Score("M")
	return M
}

func (d Dude) ScoreA() int {
	D := d.Score("D")
	return D
}

// Score builds grade according to the position.
func (d Dude) Score(p string) int {
	var v int
	for _, s := range d.Results {

		if s.Position == p {
			v += s.Score
		}
	}
	return v
}

// String implements the fmt.Stringer interface. It displays the name of the dude, it's current grades, and the number of fights he did.
func (d Dude) String() string {
	raw, _ := json.MarshallIndent(d, "", "    ")
	return string(raw)
}
