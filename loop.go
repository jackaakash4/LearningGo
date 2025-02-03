package main

import "fmt"

func main() {
    
    t1 := 1
    t2 := 1

    for i:= 1; i<10; i++ {
        fmt.Println(t1)
        nextTerm := t1 + t2
        t1 = t2
        t2 = nextTerm
    }
}

