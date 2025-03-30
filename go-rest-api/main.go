package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "log"
    "io"
)

func main(){
    
    //creating a custom request
    req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users/1", nil)

    if err != nil {
        log.Fatal("Error in creating a custom request: ", err)
    }

    //We can also add the header

    req.Header.Set("Content-Type", "application/json")

    //Executing the custom request

    resp, err := http.DefaultClient.Do(req)

    if err != nil {
        log.Fatal("Error in responsing from server: ", err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    
    if err != nil {
        log.Fatal("Error in converting response to body", err)
    }
    
    fmt.Println("response from server", string(body))

    var result struct{
            id int
            name string
            email string
        }
    

        
    if err := json.Unmarshal(body, &result); err != nil {
      log.Fatal(err)
    }
    
    fmt.Println("Response body is : ", result.id, result.name, result.email)
}
