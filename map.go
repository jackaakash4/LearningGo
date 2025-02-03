//Maps are key-value pairs. They are useful for storing and retrieving data efficiently

package main

import "fmt"

func main() {
    
    //creating map
    user := map[string]string{
        "name": "Jack",
        "email": "jackaakash44@gmail.com",
    }

    //adding a new key-value pair
    user["age"] = "22"

    //Accessing a value
    fmt.Println("User's name is ",user["name"])

    //check if a key exist
    if email, ok := user["email"]; ok {
        fmt.Println("Email: ", email)
    }
    fmt.Println(user)

    data := map[string]float32{
        "pi" : 3.14,
        "g" : 9.80,
    }
    fmt.Println(data)
    if g, ok := data["g"]; ok{
        fmt.Println("The value of acceleration due to gravity is ", g)
    }

}
