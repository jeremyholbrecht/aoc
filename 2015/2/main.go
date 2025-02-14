package main

import (
    "fmt"
    "sort"
)

func main(){
    fmt.Println(calc_amount_wrapping_paper([]int{2,3,4}))
}


// TODO: use data from the puzzle input file
func calc_amount_wrapping_paper(dimensions []int) int{
    // 2*l*w + 2*w*h + 2*h*l
    area := 2 * dimensions[0] * dimensions[1] + 2 * dimensions[1] * dimensions[2] + 2 * dimensions[0] * dimensions[2]

    // calc area shortest side
    // sorts in ascending order
    sort.Ints(dimensions)
    slack := dimensions[0] * dimensions[1]

    total := area + slack

    return total
}
