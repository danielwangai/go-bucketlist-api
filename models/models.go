package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

var (
	db database
)

// interface of functions to be used
type database interface {
	Create(interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	AutoMigrate(values ...interface{}) *gorm.DB
}

func init() {
	dotenv_err := godotenv.Load() // loads env variables from .env file
	if dotenv_err != nil {
		log.Fatal(dotenv_err)
	}
	db = Connect()
	db.AutoMigrate(Bucketlist{}, Item{})
}

// models
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
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")
	SslMode := os.Getenv("SSL_MODE")
	db, dbErr := gorm.Open("postgres", "host="+DbHost+" port=5432 dbname="+DbName+" sslmode="+SslMode)
	if dbErr != nil {
		fmt.Println("DB Connection ERROR")
		log.Fatal(dbErr)
		return nil
	}
	return db
}

func CreateBucketlist(name, description string) (*Bucketlist, error) {
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

func FetchBucketlists() (*[]Bucketlist, error) {
	var bucketlists []Bucketlist
	db.Find(&bucketlists)
	if len(bucketlists) > 0 {
		return &bucketlists, nil
	}
	return nil, errors.New("You have no bucketlists.")
}

func FetchOneBucketlist(id string) (*Bucketlist, error) {
	var bucketlist Bucketlist
	db.Where("id = ?", id).First(&bucketlist)
	if bucketlist.ID == id {
		return &bucketlist, nil
	}
	return nil, errors.New("Update unsuccessful. Bucketlist matching given id not found.")
}

func UpdateBucketlist(id, name, description string) (*Bucketlist, error) {
	// find the bucketlist by ID
	bucketlist, err := FetchOneBucketlist(id)
	if err != nil {
		return nil, errors.New("Bucketlist not found.")
	}
	// update bucketlist
	bucketlist.Name = name
	bucketlist.Description = description
	db.Save(*bucketlist)
	return *&bucketlist, nil
}

func DeleteBucketlist(id string) error {
	bucketlist, err := FetchOneBucketlist(id)
	if err != nil {
		return errors.New("Delete unsuccessful. Bucketlist not found.")
	}
	db.Delete(&bucketlist)
	return nil
}

func CreateItem(bucketlist Bucketlist, description string) (*Item, error) {
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

func FetchBucketlistItems(bucketlist Bucketlist) ([]Item, error) {
	items := bucketlist.Items
	if len(items) > 0 {
		return items, nil
	}
	return nil, errors.New("The bucketlist has no items.")
}

func FetchOneItem(id string) (*Item, error) {
	var item Item
	db.Where("id = ?", id).First(&item)
	if item.ID == id {
		return &item, nil
	}
	return nil, errors.New("The item matching id does not exist.")
}

func UpdateItem(id, description string) (*Item, error) {
	item, err := FetchOneItem(id)
	if err != nil {
		return nil, errors.New("Item matching ID not found.")
	}
	item.Description = description
	db.Save(&item)
	return *&item, nil
}

func DeleteItem(id string) error {
	item, err := FetchOneItem(id)
	if err != nil {
		return errors.New("Item matching ID not found.")
	}
	db.Delete(&item)
	return nil
}
