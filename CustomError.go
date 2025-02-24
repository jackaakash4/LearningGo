package main

import "fmt"

type validationError struct {
    Field string
    Message string
}

func (e *validationError) Error() string{
    return fmt.Sprintf("Validation Error on field: %s: %s", e.Field, e.Message)
}
//custom error 
func validateUser(name string, age int) error{
    if name == "" {
        return &validationError{
            Field: "name",
            Message: "name cannot be empty",
        }
    }
    if age < 0  {
        return &validationError{
            Field: "age",
            Message: "age should be greater than 0",
        }
    }
    return nil
}

func main(){
    err := validateUser("", -30)

    if err != nil {
        fmt.Println("Error: ", err)
    }
}

