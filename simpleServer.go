package main

import (
    "fmt"
    "flag"
    "net/http"
)

func main(){
    addr := flag.String("addr", "8080", "server port")
    flag.Parse()

    http.HandleFunc("/", home)
    http.ListenAndServe(":"+ *addr, nil)
}

func home(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "HelloWorld")
}

    
