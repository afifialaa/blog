package main

import (
	"fmt"
	"net/http"

	"github.com/afifialaa/events/database"
	"github.com/afifialaa/events/handlers"
	gorillaHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.Seed()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/blog", handlers.ReadBlog).Methods("GET")
	r.HandleFunc("/blog", handlers.CreateBlog).Methods("POST")
	r.HandleFunc("/blog", handlers.DeleteBlog).Methods("DELETE")
	r.HandleFunc("/blog", handlers.UpdateBlog).Methods("PUT")

	// Listen and serve
	fmt.Println("server is running")
	http.ListenAndServe(":8000", gorillaHandler.CORS()(r))
}
