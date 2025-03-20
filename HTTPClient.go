package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
)

func main(){

    resp, err := http.Get("https://example.com")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    fmt.Println("Response body: ", body)

}
