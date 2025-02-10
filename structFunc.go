package main
import "fmt"

type Person struct{
    Name string
    Age int
}

func (p Person) Introduce() {
    fmt.Printf("Hi, I'm %s and i'm %d years old", p.Name, p.Age)
}

func main(){
    person := Person{Name:"Aakash", Age:22}
    
    person.Introduce()
}

