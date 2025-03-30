package main

import (
    "fmt"
    "net/http"
    "io"
    "log"
)

func main(){
    

    //http client

    resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

    if err != nil {
        log.Fatal("Error in response:", err)
    }

    defer resp.Body.Close()

    //reading the response

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        log.Fatal("Error in reading the respoonse body: ", err)
    }

    fmt.Println("Response from the server is ", string(body))
}

