package main
import (
    "fmt"
)

func main(){

            assessmentMarks:=[3]int{45, 40, 38} 
            var graceMarks int =5
            finalMarks:=assessmentMarks
    
            for i := 0; i < len(finalMarks); i++ {
                   
            	finalMarks[i]+=graceMarks
                  
            }
            fmt.Println(assessmentMarks)
            fmt.Println(finalMarks)
}
