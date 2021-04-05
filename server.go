package main

import (
	"fmt"
	"net/http"

	"github.com/afifialaa/blog/config"
	"github.com/afifialaa/blog/database"
	"github.com/afifialaa/blog/handlers"
	gorillaHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Set environment variables
	config.SetEnv()

	database.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/blog", handlers.FetchArticles).Methods("GET")
	r.HandleFunc("/blog", handlers.CreateArticle).Methods("POST")
	r.HandleFunc("/blog", handlers.DeleteArticle).Methods("DELETE")
	r.HandleFunc("/blog", handlers.UpdateArticle).Methods("PUT")

	r.HandleFunc("/comment", handlers.FetchComments).Methods("GET")
	r.HandleFunc("/comment", handlers.PostComment).Methods("POST")
	r.HandleFunc("/comment", handlers.DeleteComment).Methods("DELETE")
	r.HandleFunc("/comment", handlers.UpdateComment).Methods("PUT")

	// Listen and serve
	fmt.Println("server is running")
	http.ListenAndServe(":8000", gorillaHandler.CORS()(r))
}
