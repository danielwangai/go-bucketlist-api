package controllers

import (
	"encoding/json"
	"fmt"
	"go_bucketlist_api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBucketlist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var bucketlist models.Bucketlist
	if err := json.NewDecoder(r.Body).Decode(&bucketlist); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		panic(err)
	}
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
	bucketlists, err := models.FetchBucketlists()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		panic(err)
	}
	RespondWithJson(w, http.StatusOK, bucketlists)
}

func GetOneBucketlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bucketlist, err := models.FetchOneBucketlist(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		panic(err.Error())
	}
	RespondWithJson(w, http.StatusOK, bucketlist)
}

func UpdateBucketlist(w http.ResponseWriter, r *http.Request) {

}

func DeleteBucketlist(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not implemented yet.")
}
