package main

import (

)

type Dude struct {
 Name string
 Grades [][]string

}

//CalcResult implements the calculation of DudeResult based on the arguments of DudeGrades
func (d Dude) CalcResult  () {
type DudeResult struct {
        A int
        M int
        D int
    }
    var i int
   for {
       if i>len(d.Grades){
           break
       }
       switch d.Grades[i][0] {
           case "A":
               DudeResult.A+=d.Grades[i][1]
               i++
           case "M":
               DudeResult.M+=d.Grades[i][1]
               i++
           case "D":
               DudeResult.D+=d.Grades[i][1]
               i++
       }
    
   }
   return DudeResult
}




    
