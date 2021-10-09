package main

import ("net/http"
	)

//this handler is used to handling all the different request that are made by the users
func Userhandling(response http.ResponseWriter, request *http.Request) {
    //createuserhandler
	if request.Method == "POST" {
		Createauserprofile(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}

 func Posthandling(response http.ResponseWriter, request *http.Request) {
	//createposthandler
 	if request.Method == "POST" {
		Createapost(response, request)
	} else {
	response.WriteHeader(http.StatusMethodNotAllowed)
	response.Write([]byte(`{ "message": "Incorrect Method" }`))
 	}
}
func GetUserhandler(response http.ResponseWriter, request *http.Request) {
//getuserhandler
	if request.Method == "GET" {
		GettheuserwithID(response, request)
 	} else {
 		response.WriteHeader(http.StatusMethodNotAllowed)
 		response.Write([]byte(`{ "message": "Incorrect Method" }`))
 	}
 }
func GetPosthandler(response http.ResponseWriter, request *http.Request) {
 	//getposthandler
 	if request.Method == "GET" {
 		GettingapostwithID(response, request)
 	} else {
 		response.WriteHeader(http.StatusMethodNotAllowed)
 		response.Write([]byte(`{ "message": "Incorrect Method" }`))
 	}
 }
    //getuserposthandler
func GetuserPostlisthandler(response http.ResponseWriter, request *http.Request) {
 	if request.Method == "GET" {
 		Getusersposts(response, request)
 	} else {
 		response.WriteHeader(http.StatusMethodNotAllowed)
 		response.Write([]byte(`{ "message": "Incorrect Method" }`))
 	}
}







