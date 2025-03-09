package main

import (
    "fmt"
    "time"
    "sync"
)

func printNumbers( wg *sync.WaitGroup){

    defer wg.Done()
    
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep( 500 * time.Millisecond)
    }
}

func main(){
    var wg sync.WaitGroup
    
    wg.Add(1)
    go printNumbers(&wg)
    
    wg.Wait() // Wait for the goroutine to finish

    fmt.Println("main function")
}


