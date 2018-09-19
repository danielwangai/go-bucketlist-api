package controllers

import (
	"encoding/json"
	"fmt"
	"go_bucketlist_api/models"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = models.Connect()
	fmt.Println("DB -> ", db)
}

func CreateBucketlist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var bucketlist models.Bucketlist
	if err := json.NewDecoder(r.Body).Decode(&bucketlist); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	b, err := models.CreateBucketlist(bucketlist.Name, bucketlist.Description, db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}

func GetAllBucketlists(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not implemented yet.")
}

func GetOneBucketlist(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not implemented yet.")
}

func UpdateBucketlist(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not implemented yet.")
}

func DeleteBucketlist(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not implemented yet.")
}
