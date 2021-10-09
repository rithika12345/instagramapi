package main

import (
"sync"
"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

var lock sync.Mutex

var Defaultskip = int64(0)


var Defaultlimit = int64(10)

var skip = Defaultskip
var limit = Defaultlimit

//creating user
type user struct {
Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
Email     string             `json:"email,omitempty" bson:"email,omitempty"`
Name      string             `json:"name,omitempty" bson:"name,omitempty"`
Password  string             `json:"password,omitempty" bson:"password,omitempty"`

}

//creating post
type post struct {
Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
Caption     string             `json:"caption,omitempty" bson:"caption,omitempty"`
URL         string             `json:"url,omitempty" bson:"url,omitempty"`
Posttime    string             `json:"posttime,omitempty" bson:"posttime,omitempty"`
}

func (usern *user) cons() {

if usern.Email == "" {
usern.Email = "defaultmail@email.com"
}
if usern.Name == "" {
usern.Name = usern.Email
}
if usern.Password== "" {
usern.Password="ABCD123"
}
}

func (obj *post) def() {
if obj.Caption == "" {
obj.Caption = "NO CAPTION"
}
if obj.Posttime == "" {
obj.Posttime = string(time.Now().Format(time.RFC3339))
}
if obj.URL == "" {
obj.URL = "https://defaultimage.com"
}

}