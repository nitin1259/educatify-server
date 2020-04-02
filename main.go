package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nitin1259/educatify-server/app"
	"github.com/nitin1259/educatify-server/controllers"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Educatify!")
}

// func main is just a function name but when used with package main, it serves as the application entry point.
func main() {
	fmt.Println("Welcome to educatify !")

	router := mux.NewRouter()
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/user/new/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login/", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/post/new/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/user/posts/", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/user/posts/{id}", controllers.GetPost).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api

	if err != nil {
		log.Fatal(err)
	}
}
