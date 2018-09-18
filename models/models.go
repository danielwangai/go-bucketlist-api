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
	DeletedAt *time.Time
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
	db.First(&bucketlist, id)
	if bucketlist.ID == id {
		return &bucketlist, nil
	}
	return nil, errors.New("Bucketlist matching given id not found.")
}
