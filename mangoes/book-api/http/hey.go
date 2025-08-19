package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/*
func handler (w http.ResponseWriter , r *http.Request){
	//When a client (like a browser or Postman) sends an HTTP request to your server:The server listens on a socket.When a connection is received, it parses the raw HTTP text.It builds a *http.Request object from that data.


	fmt.Fprintln(w, "from the go web server")

}

func handlerAbout ( w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"hello from about section")
}
*/

func main (){
	/*http.HandleFunc("/", handler)
	http.HandleFunc("/about", handlerAbout)
	fmt.Println("server runningat http://locahost:8080")
	http.ListenAndServe(":8080", nil)
	*/

	r := chi.NewRouter()
	r.Use(middleware.Logger)       // Logs requests
    r.Use(middleware.Recoverer)  
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Welcome to the home page powered by Chi!")
	})

	//another  route
	r.Get("/about", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "ABout chi")
	})

	//another  route
	r.Get("/contact", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Contact chi")
	})

	//dynamic routing

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request){
		userId := chi.URLParam(r, "id")
		fmt.Fprintln(w, "User Id: %s \n", userId)

	})

	//custom not found handler
	r.NotFound(func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "404 buddy, page not found", http.StatusNotFound)
	})

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))

}