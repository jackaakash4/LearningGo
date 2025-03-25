package main

import (
    "fmt"
    "net/http"
    "log"
)

func main(){

    //creating a custom request

    req, err := http.NewRequest("GET", "https://api.example.com/users", nil)

    if err != nil {
        log.Fatal(err)
    }

    //Adding header in our custom request
    
    req.Header.Add("Authorization", "Bearer token123")
    req.Header.Set("Content-Type", "application/json")

    //Add query parameter

    q := req.URL.Query()

    q.Add("page", "1")
    q.Add("limit", "10")

    req.URL.RawQuery = q.Encode()

    //Execute request with the default client

    resp, err := http.DefaultClient.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Response from custom request: ", resp)
    

}

