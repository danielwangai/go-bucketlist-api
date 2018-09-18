package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go_bucketlist_api/controllers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	dotenv_err := godotenv.Load() // loads env variables from .env file
	if dotenv_err != nil {
		log.Fatal(dotenv_err)
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	SSL_MODE := os.Getenv("SSL_MODE")
	db, db_err := gorm.Open("postgres", "host="+DB_HOST+" port=5432 dbname="+DB_NAME+" sslmode="+SSL_MODE)
	if db_err != nil {
		fmt.Println("DB Connection ERROR")
		log.Fatal(db_err)
	}
	_ = db

	router := mux.NewRouter()
	// bucketlist routes
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
