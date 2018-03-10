package riley

import ()

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

func (d Dude) A() int {
	return d.score("A")
}

func (d Dude) M() int {
	return d.score("M")
}

func (d Dude) D() int {
	return d.score("D")
}

// score builds grade according to the position.
func (d Dude) score(p string) int {
	var v int
	for _, s := range d.Results {

		if s.Position == p {
			v += s.Score
		}
	}
	return v
}
