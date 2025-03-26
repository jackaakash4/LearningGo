package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
)

func main() {

    resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

    if err != nil {
        log.Fatal("Fetching error: ", err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    fmt.Println("Response body: ", string(body))
}

