package main

import "fmt"

func main(){
    //creating a slice
    numbers := []int{10, 20, 30, 40}

    //append to a slice
    numbers = append(numbers, 60)

    //Slice a slice(get a subset
    subset := numbers[1:4]   //[20, 30, 40] index 1 is included and index 4 is excluded

    fmt.Println(numbers)
    fmt.Println(subset)

}

//slice are dynamic and flexible views into arrat
//they are commonly used than arrays bcoz they can grow and //shrink

