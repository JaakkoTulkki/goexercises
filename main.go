package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"flag"
)


func validate_name(name string) (bool, string) {
	str_len := len(name)
	if str_len == 0{
		return false, "Name not present"
	}

	if str_len < 2{
		return false, "Name needs to be greater than one character"
	}
	return true, ""
}

// Docs
// https://golang.org/pkg/net/http
// https://golang.org/pkg/io/#Writer

type Reponse struct {
    Code   int    `json:"code"`
    Result string `json:"result"`
}


func create_response(http_status int, message string) Reponse{
	response := Reponse{
		Code: http_status,
		Result: fmt.Sprintf("%s", message),
	}
	return response
}

// This is our function we are going to use to handle the request
// All handlers need to accept two arguments
// 1. Is the ResponseWriter interface, we use this to write a reponse back to the client
// 2. Is the Reponse struct which holds useful information about the request headers, method, url etc
func hello(w http.ResponseWriter, r *http.Request) {
	// We use the standard libaries WriteStirng function to write back to the ResponseWriter interface
	// See docs above
	//io.WriteString(w, fmt.Sprintf("%s %s", "hello", r.FormValue("name")))

	name := r.FormValue("name")
	is_valid, msg := validate_name(name)
	var res Reponse
	var http_status int

	if !is_valid{
		http_status = 400
	} else {
		http_status = 200
		msg = "Result: " + name
	}
	res = create_response(http_status, msg)

	json, err := json.Marshal(res)
    	if err != nil {
		log.Fatal(err)
    	}
    	w.Header().Set("Content-Type", "application/json")
    	w.Write(json)
}

func main() {

	var used_port string
	flag.StringVar(&used_port, "port", "8000", "Port to be listened")
	flag.Parse()
	used_port = ":" + used_port

	// Add ads the function thats going to handle that response
	http.HandleFunc("/", hello)
	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil
	http.ListenAndServe(used_port, nil)
}