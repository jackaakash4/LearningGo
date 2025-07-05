package main

import (
    "fmt"
)
func sliceToChannel(nums []int) <- chan int{
    out := make(chan int)
    
    go func() {
        for _, n := range nums{
            out <- n
        }
        close(out)
    }()
    return out

}
func sq(in <- chan int) <- chan int{
    out := make(chan int)
    go func(){
        for n := range in{
            out <- n*n
        }
        close(out)
    }()
    return out
}


func main(){
    //input
    nums := []int{1, 2, 3, 4, 5}
    //stage1
    dataChannel := sliceToChannel(nums)
    //stage2
    finalChannel := sq(dataChannel)

    //stage3
    for n := range finalChannel {
        fmt.Println(n)
}
}
