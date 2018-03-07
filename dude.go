package main

import (

)

type Dude struct {
 Name string
 Results []DudeResult
}

type DudeResult struct{
        Position string
        Score int
    }

const(
    V = 3
    T = 2
    D = 1
)

func (d Dude) Sum() int {
    var v int
    for _, s := range d.Results {
            v+= s.Score
        }
        return v
}




    
