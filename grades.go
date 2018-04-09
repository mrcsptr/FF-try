package riley

import(
    "encoding/json"
) 

type Grades struct{
    vals map[string]int
}

func NewGrades() *Grades {
    return &Grades{vals: make(map[string]int)}
}

func (g *Grades) Scores() map[string]int {
    m := make(map[string]int, len(g.vals))
    for k, v := range g.vals {
        m[k] = v
    }
    return m
}

func (g *Grades) AddScore(pos string, dudes ...Dude) {
	for _, d := range dudes {
		g.addScore(pos, d)
	}
}

// addScore calculates the grades of a team and returns a TeamStat containing them
func (g *Grades) addScore(position string, d Dude) {

	switch position {
	case "A":
		g.vals["A"] += d.A()
	case "M":
		g.vals["M"] += d.M()
	case "D":
		g.vals["D"] += d.D()
	}
}


// String returns the json representation of a dude.
func (g *Grades) String() string {
	raw, _ := json.Marshal(g.vals)
	return string(raw)
}


var Positions = []string{"A", "M", "D"}
