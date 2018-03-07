package riley

// Dude a type holding its characteristics.
type Dude struct {
	Name    string
	Results []DudeResult
}

// calculating the grades
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
