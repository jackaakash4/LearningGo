package main

import "fmt"

func main() {
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is not even");
    }

    num := 9

    if num < 0 {
        fmt.Println(num + "is positive")
    } else if num > 0 {
        fmt.Println(num + "is negative")
    } else {
        fmt.Println(num + " is zero")
    }
}

