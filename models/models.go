package models

import (
	"errors"
	"fmt"
	"log"
	"os"
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

// DB connection
func Connect() (db *gorm.DB) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	SSL_MODE := os.Getenv("SSL_MODE")
	db, db_err := gorm.Open("postgres", "host="+DB_HOST+" port=5432 dbname="+DB_NAME+" sslmode="+SSL_MODE)
	if db_err != nil {
		fmt.Println("DB Connection ERROR")
		log.Fatal(db_err)
		return nil
	}
	return db
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

func FetchBucketlistItems(bucketlist Bucketlist, db *gorm.DB) ([]Item, error) {
	items := bucketlist.Items
	if len(items) > 0 {
		return items, nil
	}
	return nil, errors.New("The bucketlist has no items.")
}

func FetchOneItem(id string, db *gorm.DB) (*Item, error) {
	var item Item
	db.Where("id = ?", id).First(&item)
	if item.ID == id {
		return &item, nil
	}
	return nil, errors.New("The item matching id does not exist.")
}

func UpdateItem(id, description string, db *gorm.DB) (*Item, error) {
	item, err := FetchOneItem(id, db)
	if err != nil {
		return nil, errors.New("Item matching ID not found.")
	}
	item.Description = description
	db.Save(&item)
	return *&item, nil
}

func DeleteItem(id string, db *gorm.DB) error {
	item, err := FetchOneItem(id, db)
	if err != nil {
		return errors.New("Item matching ID not found.")
	}
	db.Delete(&item)
	return nil
}
