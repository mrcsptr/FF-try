package main

import (

)
// Dude a type containing the name of the dude and his collection of results
type Dude struct {
 Name string
 Results []DudeResult
}

// DudeResult contains one result, defining the position occupied then by the player, and the score obtained.
type DudeResult struct{
        Position string
        Score int
}

// Definition of the values for a win (V), tie (T), and a loss (L) (unused for now).
const(
    W = 3
    T = 2
    L = 1
)
// Sum produces the current grades for the player, depending on the position that will be occupied: Attack (A), Middle (M), and Defense (D).
func (du Dude) Sum() (int, int, int){
    var a int
    var m int
    var d int
    for _, s := range du.Results {
        switch s.Position {
            case "A" :
                a+= s.Score
            case "M" :
                m+= s.Score
            case "D" :
                d+= s.Score
        }
    }
    return a, m, d
}




    
