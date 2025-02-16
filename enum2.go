package main

import "fmt"

type Status struct {
    Code int
    Message string
}

const (
        Pending Status = Status{0, "Pending")
        Approved Status = Status{1, "Approved")
        Rejected Status = Status{2, "Rejected")
)

func main() {
    status := Pending

    fmt.Println("Your status is ", status.Message)
}
