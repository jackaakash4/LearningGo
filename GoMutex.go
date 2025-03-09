//A sync.Mutex
package main

import (
    "fmt"
    "sync"
)

var counter int
var mu sync.Mutex

func increment( wg *sync.WaitGroup) {
    defer wg.Done()

    for i := 0; i < 10; i++ {
        //we are locking counter so that it can only be accessed by one goroutine. Counter is a shared resource
        mu.Lock()
        counter++
        mu.Unlock() //unlock the mutex

        }
    }

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go increment(&wg)
    go increment(&wg)
    wg.Wait()
    fmt.Println("counter: ", counter)
}


