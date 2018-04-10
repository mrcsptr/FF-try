package riley

import(
    "encoding/json"
) 

// Grades holds the grades of all teams
type Grades struct{
    vals map[string]int
}

// NewGrades sends the calculated grades into the map contained in Grades
func NewGrades() *Grades {
    return &Grades{vals: make(map[string]int)}
}

// Scores initialyzes the map, clone it and return the clone to avoid bad things
func (g *Grades) Scores() map[string]int {
    m := make(map[string]int, len(g.vals))
    for k, v := range g.vals {
        m[k] = v
    }
    return m
}

// AddScore calculates all the grades for a team
func (g *Grades) AddScore(pos string, dudes ...Dude) {
	for _, d := range dudes {
		g.addScore(pos, d)
	}
}

// addScore calculates the grade of a player of the team according to the position he is within the team and add it to the current grade value of the team.
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
