package main

import (
    "fmt"
    "encoding/json"
    "log"
    )

func main() {

    data := map[string]string{
        "name": "Jack",
        "email": "jack@gmail.com",
    }

    fmt.Println("Data: ", data)

    //encoding in json

    jsonData, err := json.Marshal(data)

    if err != nil {
        log.Fatal("Error in encoding in json:", err)
    }

    fmt.Println("Encoded data: ", jsonData)
    fmt.Println("Encoded data in string: ", string(jsonData))


    //now unmarshaling the json data

    var result struct{
        name string
        email string
    }

    if err := json.Unmarshal(jsonData, &result); err != nil {
        log.Fatal(err)
        }

    fmt.Printf("User's name is %s and email is %s", result.name, result.email)

}
