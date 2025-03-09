package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Hello"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "World"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg := <- ch1:
            fmt.Println(msg)
        case msg := <- ch2:
            fmt.Println(msg)
        }
    }
}
