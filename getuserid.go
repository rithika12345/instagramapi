package main

import (
"context"
"encoding/json"
"errors"
"fmt"
"net/http"
"path"
"time"

"go.mongodb.org/mongo-driver/bson/primitive"
)

//retuning the user id back to the function
func gettinguserwithId(id primitive.ObjectID) (user, error) {
		var users user
		collection := client.Database("Appointy").Collection("Instagramusers")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		err := collection.FindOne(ctx, user{Id: id}).Decode(&users)
		if users.Id != id {
		err = errors.New("Error 400: ID not present")
		}
		return users, err
}

//getting a user with a specific id
func GettheuserwithID(response http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
		}
		response.Header().Set("content-type", "application/json")
		fmt.Println(path.Base(request.URL.Path))
		id, _ := primitive.ObjectIDFromHex(path.Base(request.URL.Path))
		getuserwithId, err := gettinguserwithId(id)
		if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
		}
		json.NewEncoder(response).Encode(getuserwithId)

}

