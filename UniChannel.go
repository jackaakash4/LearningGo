package main

import "fmt"

func sendData(ch chan int){
    ch <- 42
}

func receiveData(ch chan int){
    receive := <- ch
    fmt.Println("Received data from channel is ", receive)
}


func main() {
    ch := make(chan int)
    go sendData(ch)

    receiveData(ch)
}
