package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "encoding/json"

)

func main() {

    //creating a custom request
    data:= map[string]string{
        "name":"Jack",
        "email":"jack@gmail.com",
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Fatal("Error in converting into json", err)
    }



    req, err := http.NewRequest("POST","https://jsonplaceholder.typicon.com/users", jsonData)
    
    if err != nil {
        log.Fatal("Error in posting the data to an api ", err)
    }

    resp, err := http.DefaultClient.Do(req)


    //Response handling

    if err != nil {
        log.Fatal("Error in responsing ", err)
    }

    defer resp.Body.Close()

    //checking status

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Request faild with status: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)

    fmt.Println("Response body", body)
}
       
