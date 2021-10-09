package main

import (
"context"
"encoding/json"
"fmt"
"net/http"
"time"
"crypto/sha256"
"go.mongodb.org/mongo-driver/bson/primitive"
)

//here we are creating a new user 
func Createauserprofile(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json")
		var users user
		_ = json.NewDecoder(request.Body).Decode(&users)
		//encrypting the password
		sha256 := sha256.Sum256([]byte(users.Password))
		users.cons()
		lock.Lock()
		defer lock.Unlock()
		collection := client.Database("Appointy").Collection("Instagramusers")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var sx32 string = fmt.Sprintf("%x", sha256)
		users.Password=sx32
		//storingthedata
		result, _ := collection.InsertOne(ctx, users)
		users.Id = result.InsertedID.(primitive.ObjectID)
		json.NewEncoder(response).Encode(users)
		fmt.Println(users)
}


