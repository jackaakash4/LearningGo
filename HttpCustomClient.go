package main

import (
    "fmt"
    "time"
    "log"
    "io"
    "net/http"
)

func main() {
    

    //creating a custom client

    client := &http.Client{
        Timeout : 10 * time.Second,
        Transport : &http.Transport{
                        MaxIdleConns: 10,
                        IdleConnTimeout: 30 * time.Second,
                    },
                }
    
    resp, err := client.Get("https://example.com")

    if err != nil {
        log.Fatal(err)
    }
    
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        log.Fatal("Error in reading response body", err)
    }

    fmt.Println("Response body: ", body)
}

    
