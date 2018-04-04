package riley

import ()

// Team a type holding its characteristics.
type Team struct {
	Name string
	Team []Teammate
}

// TeamStat a type holding a team's name and its grades for a fight
type TeamStat struct {
	Name   string
	Grades []int
}

// scoresum calculates the grades of a team and returns a TeamStat containing them
func (d []riley.Dude, t Team) scoresum() TeamStat {
	for _, team := range t.Team {
		var Agrade, Mgrade, Dgrade int

		for _, dudes := range d {
			if t.Team.Name == d.Name {
				switch t.Teammate.Position {
				case "A":
					gradeA := riley.A(d)
				case "M":
					gradeM := riley.M(d)
				case "D":
					gradeD := riley.D(d)
				}
				break
			}

		}
		Agrade += gradeA
		Mgrade += gradeM
		Dgrade += gradeD

	}
	var s TeamStat
	s.Name = t.Name
	s.Grades = []int{Agrade, Mgrade, Dgrade}
	return s
}
