package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BaseModel struct {
	ID        string `gorm:"primary_key;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []Item
}

type Bucketlist struct {
	BaseModel
	Name        string `gorm:"not null"`
	Description string `gorm:"not null;size=400"`
}

type Item struct {
	BaseModel
	Description  string `gorm:"not null;size:200"`
	Bucketlist   Bucketlist
	BucketlistId string
}

func CreateBucketlist(name, description string, db *gorm.DB) (*Bucketlist, error) {
	var bucketlist Bucketlist
	bucketlist.ID = uuid.Must(uuid.NewV4()).String()
	bucketlist.Name = name
	bucketlist.Description = description
	db.Create(&bucketlist)
	if len(bucketlist.ID) > 0 {
		return &bucketlist, nil
	}
	return nil, errors.New("An error occured. Create operation unsuccessful.")
}

func FetchBucketlists(db *gorm.DB) (*[]Bucketlist, error) {
	var bucketlists []Bucketlist
	db.Find(&bucketlists)
	if len(bucketlists) > 0 {
		return &bucketlists, nil
	}
	return nil, errors.New("You have no bucketlists.")
}

func FetchOneBucketlist(id string, db *gorm.DB) (*Bucketlist, error) {
	var bucketlist Bucketlist
	db.Where("id = ?", id).First(&bucketlist)
	if bucketlist.ID == id {
		return &bucketlist, nil
	}
	return nil, errors.New("Update unsuccessful. Bucketlist matching given id not found.")
}

func UpdateBucketlist(id, name, description string, db *gorm.DB) (*Bucketlist, error) {
	// find the bucketlist by ID
	bucketlist, err := FetchOneBucketlist(id, db)
	if err != nil {
		return nil, errors.New("Bucketlist not found.")
	}
	// update bucketlist
	bucketlist.Name = name
	bucketlist.Description = description
	db.Save(*bucketlist)
	return *&bucketlist, nil
}

func DeleteBucketlist(id string, db *gorm.DB) error {
	bucketlist, err := FetchOneBucketlist(id, db)
	if err != nil {
		return errors.New("Delete unsuccessful. Bucketlist not found.")
	}
	db.Delete(&bucketlist)
	return nil
}

func CreateItem(bucketlist Bucketlist, description string, db *gorm.DB) (*Item, error) {
	var item Item
	item.ID = uuid.Must(uuid.NewV4()).String()
	item.Description = description
	item.BucketlistId = bucketlist.ID
	db.Create(&item)
	if len(item.ID) > 0 {
		return &item, nil
	}
	return nil, errors.New("Item could not be created.")
}
