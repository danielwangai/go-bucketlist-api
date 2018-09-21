package main

import (
	"log"
	"net/http"

	"go_bucketlist_api/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	dotenv_err := godotenv.Load() // loads env variables from .env file
	if dotenv_err != nil {
		log.Fatal(dotenv_err)
	}
}

func main() {
	// bucketlist routes
	router := mux.NewRouter()
	router.HandleFunc("/bucketlists", controllers.CreateBucketlist).Methods("POST")
	router.HandleFunc("/bucketlists", controllers.GetAllBucketlists).Methods("GET")
	router.HandleFunc("/bucketlists/{id}", controllers.GetOneBucketlist).Methods("GET")
	router.HandleFunc("/bucketlists/{id}", controllers.UpdateBucketlist).Methods("PUT")
	router.HandleFunc("/bucketlists/{id}", controllers.DeleteBucketlist).Methods("DELETE")
	// bucketlist item routes
	router.HandleFunc("/bucketlists/{id}/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/bucketlists/{id}/items", controllers.GetBucketlistItems).Methods("GET")
	router.HandleFunc("/bucketlists/{id}/items/{itemId}", controllers.GetOneItem).Methods("GET")
	router.HandleFunc("/bucketlists/{id}/items/{itemId}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/bucketlists/{id}/items/{itemId}", controllers.DeleteItem).Methods("DELETE")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
