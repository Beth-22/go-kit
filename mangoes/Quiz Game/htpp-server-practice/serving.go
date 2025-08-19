package main 


import (
	//"fmt"
	"net/http"
	"htpp-server-practice/api"
)

/*
//to implement an interface on it, http.Handler interface
type helloHandler struct{}


//ServeHttp is a method in http.Handler interface

func( h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello from serverHttp method")
}

func main(){
	
    handler := helloHandler{}
	//a new value of the struct, saved to "handler" variable-an instance of ur custom type. If your struct had fields, like:type helloHandler struct {Greeting string} Then you could initialize it like: handler := helloHandler{Greeting: "Hi there!"}



	//start server
	fmt.Println("starting server on :8080")
	http.ListenAndServe(":8080", handler)
	// routes all request to handler

	
}*/



func main(){

	srv :=api.NewServer()
	http.ListenAndServe(":8080", srv)

	
}