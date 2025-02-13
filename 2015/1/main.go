package main

import "fmt"

const instructions  = ")())())"

func main(){
    fmt.Println(calc_floor(instructions))
}

func calc_floor(parentheses string) int{
    floor := 0

    for _, parenthesis := range parentheses{
        if parenthesis  == '('{
            floor += 1
        } else {
            floor -= 1
        }
    }

    return floor
}