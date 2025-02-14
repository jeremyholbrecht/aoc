package main

import (
    "fmt"
    "sort"
    "os"
    "log"
    "bufio"
)

const PATH = "data.txt"

func main(){
    fmt.Println(calc_amount_wrapping_paper([]int{2,3,4}))

    //TODO: convert each text line to a int slice
    for _, line := range get_text(PATH){
        fmt.Println(line)
    }
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

func get_text(path string) []string{
    file, err := os.Open(PATH)

    if err != nil {
        log.Fatal(err)
    }

   scanner := bufio.NewScanner(file)
   scanner.Split(bufio.ScanLines)
   var text []string

   for scanner.Scan(){
      text = append(text, scanner.Text())
   }

   file.Close()

   return text
}
