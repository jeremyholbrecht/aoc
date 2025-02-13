package main

import (
    "fmt"
    "os"
    "log"
)

const path = "data.txt"

func main(){
   // fmt.Println(calc_floor(path))
   find_position_basement(path)
}

//part 1
func calc_floor(path string)int{
    floor := 0
    parentheses, err := os.ReadFile(path)

    if err != nil {
        log.Fatal(err)
    }

    for _, parenthesis := range parentheses{
        if parenthesis == '('{
            floor += 1
        } else {
            floor -= 1
        }
    }
    return floor
}

//part 2
func find_position_basement(path string){
    floor := 0
    parentheses, err := os.ReadFile(path)

    if err != nil {
        log.Fatal(err)
    }

    for i, parenthesis := range parentheses{
        if parenthesis == '('{
            floor += 1
        } else {
            floor -= 1
        }

        // we need to add 1 to the index because in this puzzle the first element is at index 1 and not 0
        if floor == -1 {
            fmt.Printf("%d basement reached \n", i + 1)

        }
    }

}