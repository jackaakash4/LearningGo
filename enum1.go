package main

import "fmt"

//create custom type

type Day int

const (
        Sunday Day = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday

    )

func (d Day) String() string{
    return [...] string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}[d]
}

func main(){
    today := Tuesday

    fmt.Println("Today is " , today)

    switch today {
        case Saturday:
            fmt.Println("It's a weekend! Relax")
        default:
           fmt.Println("It's a weekday")
        }
    }

