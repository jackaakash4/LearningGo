package main

import "fmt"

func produce(n ...int) <-chan int{
    out := make(chan int)
    
   go func(){
    for _, i := range n {
        out <- i
    }
    close(out)
}()
    return out
}

func sqr(num <-chan int) <-chan int{
    out := make(chan int)
    go func(){
    for i:= range num {
        out <- i * i
    }
    close(out)
}()
    return out
}



func main() {
    
    nums := produce(1, 2, 3, 4)
    squared := sqr(nums)

    for n := range squared {
        fmt.Println(n)
    }
}
