package controllers

import (
	"encoding/json"
	"go_bucketlist_api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBucketlist(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()
	var bucketlist models.Bucketlist
	if err := json.NewDecoder(r.Body).Decode(&bucketlist); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	if len(bucketlist.Name) == 0 || len(bucketlist.Description) == 0 {
		RespondWithError(w, http.StatusBadRequest, "Bucketlist name and description required.")
		return
	}
	b, err := models.CreateBucketlist(bucketlist.Name, bucketlist.Description)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Something went wrong while creating the bucketlist.")
		return
	}
	RespondWithJson(w, http.StatusCreated, b)
}

func GetAllBucketlists(w http.ResponseWriter, r *http.Request) {
	bucketlists, err := models.FetchBucketlists()
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, bucketlists)
}

func GetOneBucketlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bucketlist, err := models.FetchOneBucketlist(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, bucketlist)
}

func UpdateBucketlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var bucketlist models.Bucketlist
	if err := json.NewDecoder(r.Body).Decode(&bucketlist); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(bucketlist.Name) == 0 || len(bucketlist.Description) == 0 {
		RespondWithError(w, http.StatusBadRequest, "Cannot update with empty fields.")
		return
	}
	b, e := models.UpdateBucketlist(params["id"], bucketlist.Name, bucketlist.Description)
	if e != nil {
		RespondWithError(w, http.StatusBadRequest, e.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, b)
}

func DeleteBucketlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bucketlist := models.DeleteBucketlist(params["id"])
	if bucketlist != nil {
		RespondWithError(w, http.StatusNotFound, "Bucketlist not found.")
		return
	}
	RespondWithJson(w, http.StatusNoContent, "Bucketlist deleted successfully.")
}
