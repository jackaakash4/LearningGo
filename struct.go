package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func main() {
    person := Person{Name:"Aakash", Age: 22}

    fmt.Println("My name is ", person.Name, " and my age is ", person.Age)

}
