package main

import "fmt"

func describe(i interface{}){
    
    switch v:= i.(type) {
    case int: 
        fmt.Printf("Integer: %v \n", v)
    case string:
        fmt.Printf("String: %v \n", v)
    case bool:
        fmt.Printf("Boolean: %v \n", v)
    default:
        fmt.Printf("Unknown type: %T \n", v)
    }
}

func main() {
    describe(42)
    describe("JackAakash")
    describe(true)
    describe(3.14)
}
