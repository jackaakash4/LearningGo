package main

import (
    "fmt"
    "flag"
    "net/http"
    "time"
)

func main(){
    addr := flag.String("addr", "8080", "Server Port")
    flag.Parse()

    http.HandleFunc("/", events)
    http.ListenAndServe(":"+*addr, nil)
}

func events(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/event-stream")

    tokens := []string{"This", "is", "produced", "by", "the", "LLM", "made", "by", "Jack"}

    for _, token := range tokens{
        content := fmt.Sprintf("data = %s \n\n", string(token))
        w.Write([]byte(content))
        w.(http.Flusher).Flush()

        time.Sleep(time.Millisecond * 240)
    }
}
