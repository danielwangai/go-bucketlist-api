package controllers

import (
	"encoding/json"
	"fmt"
	"go_bucketlist_api/models"
	"net/http"
)

func CreateBucketlist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var bucketlist models.Bucketlist
	if err := json.NewDecoder(r.Body).Decode(&bucketlist); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		panic(err)
	}
	fmt.Println("DECODE - ", bucketlist.Name)
	if len(bucketlist.Name) == 0 || len(bucketlist.Description) == 0 {
		RespondWithError(w, http.StatusBadRequest, "Bucketlist name and description required.")
		panic("Bucketlist name and description required.")
	}
	b, err := models.CreateBucketlist(bucketlist.Name, bucketlist.Description)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Something went wrong while creating the bucketlist.")
		panic(err)
	}
	RespondWithJson(w, http.StatusCreated, b)
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
