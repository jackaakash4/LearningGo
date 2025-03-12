package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int) {
    for job := range jobs {
        fmt.Println("Worker:", id, "Started job", job)
        time.Sleep(time.Second)
        fmt.Println("Worker:", id, "finished job:", job)
    }
}



func main(){

    jobs := make(chan int, 10)

    //start 3 worker
    for w := 1; w <= 3; w++ {
        go worker(w, jobs)
    }

    //send 5 jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

}

