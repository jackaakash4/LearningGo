package main

import "fmt"

func main() {
    
    numbers := []int {10, 20, 30, 40, 50}

    //append to a slice
    //append add data to end of the slice

    numbers = append(numbers, 60)
    fmt.Println(numbers)    
    
    //slice a slice(to get subset)
    subset := numbers[1:4]

    fmt.Println(subset)
}



