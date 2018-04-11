package filesystem

import (
	"fmt"
    
    "github.com/moxar/riley"
)

// Combat gets the teams that will fight
func (g riley.Grades) Combat(TeamA, TeamB string) {

	switch {
	case g["TeamA"].A > g["TeamB"].D:
		round1 := 2
	case g["TeamA"].A == g["TeamB"].D:
		round1 := 1
	case g["TeamA"].A < g["TeamB"].D:
		round1 := 0
	}
	switch {
	case g["TeamA"].M > g["TeamB"].M:
		round2 := 2
	case g["TeamA"].M == g["TeamB"].M:
		round2 := 1
	case g["TeamA"].M < g["TeamB"].M:
		round2 := 0
	}
	switch {
	case g["TeamA"].D > g["TeamB"].A:
		round3 := 2
	case g["TeamA"].D == g["TeamB"].A:
		round3 := 1
	case g["TeamA"].A < g["TeamB"].A:
		round3 := 0
	}
	result := round1 + round2 + round3
	switch {
	case result > 3:
		fmt.Println(TeamA, " wins !")
	case result == 3:
		fmt.Println("It's a draw between ", TeamA, " and ", TeamB, " !")
	case result < 3:
		fmt.Println(TeamB, " wins !")
	}

}
