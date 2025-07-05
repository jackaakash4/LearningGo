package main

import (
    "fmt"
    "time"
)

func doWork(done <- chan bool){
    for {
        select{
        case <- done:
            return
        default:
            fmt.Println("Doing work")
        }
}
}

func main(){
    done := make(chan bool)
    
    go doWork(done)

    time.Sleep(2 * time.Second)
    
    close(done)
}
