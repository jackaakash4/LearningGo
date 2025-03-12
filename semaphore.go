package main

import (
    "fmt"
    "time"
    "sync"
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup){
    
    defer wg.Done()
    sem <- struct{}{}  //acquire the semaphore
    fmt.Printf("Worker %d starded ", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d finished", id)
    
    <-sem
}

func main() {
    const numWorker = 5
    const maxConcurrent = 2

    sem := make(chan struct{}, maxConcurrent)
    
    var wg sync.WaitGroup

    for i := 1; i <= numWorker; i++ {
        wg.Add(1)
        go worker(i, sem, &wg)
    }
    wg.Wait()
}



