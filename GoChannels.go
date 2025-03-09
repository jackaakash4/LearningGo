package main

import (
    "fmt"
    "time"
)

func printNumber( ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- 1  //send data to the channel
        time.Sleep( 500 * time.Millisecond)
    }
    close(ch) //close the channel when done
}

func main() {
    ch := make(chan int) 
    go printNumber(ch)
     
    for num := range ch { //Receive data from channel
        fmt.Println(num)
    }
    fmt.Println("Main Function")
}
