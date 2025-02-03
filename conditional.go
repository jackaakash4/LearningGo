package main

import "fmt"

func main() {
    num := 10

    if num > 0 {
        fmt.Println(num, " is positive")
    } else if num < 0 {
        fmt.Println(num, " is negative number")
    } else {
        fmt.Println(num, " is zero")
    }
}

