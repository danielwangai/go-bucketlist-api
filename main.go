package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go_bucketlist_api/controllers"
	"go_bucketlist_api/models"

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
	db.AutoMigrate(models.Bucketlist{}, models.Item{})
	// b, _ := models.CreateBucketlist("name 3", "description 2", db)
	// fmt.Println(b)
	// fmt.Println(models.UpdateBucketlist("e1683d09-20b8-40a9-a29f-63ccd7a42794", "name", "descr", db))
	// fmt.Println(models.FetchOneBucketlist("e1683d09-20b8-40a9-a29f-63ccd7a42794", db))
	d, _ := models.FetchOneBucketlist("ec55dfd7-b670-40ed-be32-fb4a777e8eb9", db)
	c, _ := models.CreateItem(*d, "descr", db)
	fmt.Println(c)
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
