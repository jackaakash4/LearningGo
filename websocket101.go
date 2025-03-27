package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websockets"
)

//upgrade http to websockets

var upgrader = websockets.Upgrader {
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,

    //Allow connections from any origin from development

    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request){
    
    conn, err := upgrader.Upgrade(w, r, nil)
    
    if err != nil {
        log.Println("Websocket upgrade error: ", err)
        return
    }
    
    defer conn.Close()

    //simple echo functionality

    for {
        //Read message from client

        messageType, message, err := conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
                log.Println("Websocket err: %v", err)
            }
            break
        }

        log.Printf("Received message: %s", message)

        //Echo message back to the client

        if err := conn.WriteMessage(messageType, message); err != nil {
            log.Prinln("Write Error: ", err)
            break
        }
    }:wq
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)

    fmt.Println("Websocket server is running on: 8080 port")

    if err := http.ListenAndServer(":8080", nil); err != nil{
        log.Println("Error in listening and serving: ", err)
    }
}

