package main

import (
    
    "context"
    "fmt"
    "net/http"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
var client *mongo.Client
func main() {
    fmt.Println("The application is running successfully")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    clientOptions := options.Client().ApplyURI("mongodb+srv://rithika:pustu100@cluster0.wt3zo.mongodb.net/test")
    client, _ = mongo.Connect(ctx, clientOptions)
    http.HandleFunc("/users", Userhandling)
    http.HandleFunc("/posts", Posthandling)
    http.HandleFunc("/users/", GetUserhandler)
    http.HandleFunc("/posts/", GetPosthandler)
    http.HandleFunc("/posts/users/", GetuserPostlisthandler)
    http.ListenAndServe(":12345", nil)
}