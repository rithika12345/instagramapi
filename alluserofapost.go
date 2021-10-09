package main

import (
"context"
"encoding/json"
"fmt"
"net/http"
"strconv"
"time"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo/options"
)

//returning a list of all the post of a user to the function
func Checkuserspost(email string) []post {
		collection := client.Database("Appointy").Collection("Instagramposts")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		opts := options.Find()
		opts.SetSort(bson.D{{Key: "url", Value: 1}})
		opts.Skip = &skip
		opts.Limit = &limit
		cursor, _ := collection.Find(ctx, bson.D{
		{Key: "user.email", Value: email},
		}, opts)
		var postsreturn []post
		var posts post
		for cursor.Next(ctx) {
		cursor.Decode(&posts)
		postsreturn = append(postsreturn, posts)
		}
		return postsreturn
}

// getting a list of all the post by a specific user
func Getusersposts(response http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
		}
		response.Header().Set("content-type", "application/json")
		fmt.Println((request.URL.Query()["user"][0]))
		if len(request.URL.Query()["limit"]) != 0 {
		limit, _ = strconv.ParseInt(request.URL.Query()["limit"][0], 0, 64)
		}
		if len(request.URL.Query()["ofset"]) != 0 {
		skip, _ = strconv.ParseInt(request.URL.Query()["offset"][0], 0, 64)
		}
		email := request.URL.Query()["user"][0]
		userposts := Checkuserspost(email)
		if len(userposts) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "No post of the user" }`))
		return
		}
		json.NewEncoder(response).Encode(userposts)
		skip = Defaultskip
		limit = Defaultlimit
}