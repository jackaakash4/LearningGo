package main

import (
    "fmt"
    "sync"
    )


func producer(nums ...int) <-chan int {
    out := make(chan int)

    go func(){
        for _, n := range nums{
            
            out <- n

        }
        close(out)
    }()

    return out
}

func square(in <-chan int) <-chan int{
    out := make(chan int)

    go func() {
        for n:= range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}


func merge(channels ...<-chan int) <-chan int{
    var wg sync.WaitGroup
    out := make(chan int)

    //Fan-in
    for _, ch := range channels {
        wg.Add(1)
        go func(ch <-chan int){
            for n := range ch{
                out <-n
            }
            wg.Done()
        }(ch)
    }
    //close the output channel when all goroutines are done

    go func(){
        wg.Wait()
        close(out)
    }()
    return out
}

func main(){
    in := producer(1, 2, 3, 4)

    //Fan-out

    ch1 := square(in)
    ch2 := square(in)
    
    //Fan-in
    for n := range merge(ch1, ch2) {
        fmt.Println(n)
    }

}
