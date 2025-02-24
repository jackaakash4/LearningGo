package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "strconv"
    "log"
    
    "github.com/gorilla/mux"
)

type User struct{
    ID int `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
}

//In-memory user database for simplicity
var users = []User{
    {ID:1, Username: "jack", Email:"jackaakash@gmail.com"},
    {ID:2, Username: "jackaakash", Email:"jack@gmail.com"},
}

// GetUsers return all users for /users

func GetUsers(w http.ResponseWriter, r *http.Request){
    
    json.NewEncoder(w).Encode(users)
    }

// GetUser returns a single user by ID

func GetUser(w http.ResponseWriter, r *http.Request){

    params := mux.Vars(r)
    
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    for _, user := range users {
        if user.ID == id {
            json.NewEncoder(w).Encode(user)
            return
        }
    }

    http.Error(w, "User not found", http.StatusBadRequest)

}

func main(){

    router := mux.NewRouter()
    
    router.HandleFunc("/users", GetUsers).Methods("GET")

    router.HandleFunc("/users/{id}", GetUser).Methods("GET")

    fmt.Println("Server is running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
