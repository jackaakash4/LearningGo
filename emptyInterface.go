package main

import "fmt"

func main(){
    printValue(40)
    printValue("JackAakash")
    printValue(3.14)
}

func printValue(v interface{} ){
    fmt.Println(v)
}

