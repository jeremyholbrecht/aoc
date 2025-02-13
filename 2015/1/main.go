package main

import (
    "fmt"
    "os"
    "log"
)

const path = "data.txt"

func main(){
    fmt.Println(calc_floor(path))
}

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