package main

import (
	"log"
	"net/http"

	"go_bucketlist_api/controllers"
	"go_bucketlist_api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// func init() {

// }

func main() {
	dotenv_err := godotenv.Load() // loads env variables from .env file
	if dotenv_err != nil {
		log.Fatal(dotenv_err)
	}
	db := models.Connect()
	if db == nil {
		log.Fatal("Connection not established")
	}
	db.AutoMigrate(models.Bucketlist{}, models.Item{})
	// b, _ := models.CreateBucketlist("name 3", "description 2", db)
	// fmt.Println(b)
	// fmt.Println(models.UpdateBucketlist("e1683d09-20b8-40a9-a29f-63ccd7a42794", "name", "descr", db))
	// fmt.Println(models.FetchOneBucketlist("e1683d09-20b8-40a9-a29f-63ccd7a42794", db))
	// d, _ := models.FetchOneBucketlist("ec55dfd7-b670-40ed-be32-fb4a777e8eb9", db)
	// c, _ := models.CreateItem(*d, "descr", db)
	// fmt.Println(c)
	// i := models.DeleteItem("67eb0bf4-65cc-464e-a6f3-503c08e656b0", db)
	// fmt.Println(i)
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
