package riley

// Dude a type holding its characteristics.
type Dude struct {
	Name    string
	Results []DudeResult
}

// A returns the grade of the dude for the position A.
func (d Dude) A() int {
	return d.score("A")
}

// see A.
func (d Dude) M() int {
	return d.score("M")
}

// see A.
func (d Dude) D() int {
	return d.score("D")
}

// score builds score according to the selected position by summing all occurrences of grades related to said position.
func (d Dude) score(p string) int {
	var v int
	for _, s := range d.Results {

		if s.Position == p {
			v += s.Score
		}
	}
	return v
}
