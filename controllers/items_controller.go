package controllers

import (
	"encoding/json"
	"go_bucketlist_api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	params := mux.Vars(r)
	bucketlist, bucketErr := models.FetchOneBucketlist(params["id"])
	if bucketErr != nil {
		RespondWithError(w, http.StatusNotFound, "Bucketlist not found.")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		RespondWithError(w, http.StatusNotFound, "Invalid request payload.")
		return
	}
	if len(item.Description) == 0 {
		RespondWithError(w, http.StatusNotFound, "Item description cannot be empty.")
		return
	}
	i, itemErr := models.CreateItem(*bucketlist, item.Description)
	if itemErr != nil {
		RespondWithError(w, http.StatusNotFound, itemErr.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, i)
}

func GetBucketlistItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bucketlist, bucketErr := models.FetchOneBucketlist(params["id"])
	if bucketErr != nil {
		RespondWithError(w, http.StatusNotFound, "Bucketlist not found.")
		return
	}
	items, itemErr := models.FetchBucketlistItems(*bucketlist)
	if itemErr != nil {
		RespondWithError(w, http.StatusNotFound, itemErr.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, items)
}

func GetOneItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item, itemErr := models.FetchOneItem(params["itemId"])
	if itemErr != nil {
		RespondWithError(w, http.StatusNotFound, itemErr.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	params := mux.Vars(r)
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(item.Description) == 0 {
		RespondWithError(w, http.StatusBadRequest, "Item description required to update")
		return
	}
	i, itemErr := models.UpdateItem(params["itemId"], item.Description)
	if itemErr != nil {
		RespondWithError(w, http.StatusNotFound, itemErr.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, i)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.DeleteItem(params["itemId"])
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, map[string]string{"message": "Item deleted successfully."})
}
