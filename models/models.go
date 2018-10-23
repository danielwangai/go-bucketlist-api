package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"

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
	db.AutoMigrate(Bucketlist{}, Item{}, User{})
}

// models
type BaseModel struct {
	ID        string `gorm:"primary_key;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bucketlist struct {
	BaseModel
	Name        string `gorm:"not null"`
	Description string `gorm:"not null;size=400"`
	Items       []Item
}

type Item struct {
	BaseModel
	Description  string `gorm:"not null;size:200"`
	Bucketlist   Bucketlist
	BucketlistId string
}

type User struct {
	BaseModel
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type Token struct {
	UserId    string
	Email     string
	AuthToken string
	jwt.StandardClaims
}

// user callback
func (user *User) BeforeCreate() error {
	mailErr := ValidateEmail(user.Email)
	passErr := ValidatePassword(user.Password)
	if mailErr != nil || passErr != nil {
		return errors.New("An error occured when creating the user.")
	}
	return nil
}

// helpers
func ValidateEmail(email string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(email) == false {
		return errors.New("Invalid email format.")
	}
	return nil
}

func ValidatePassword(password string) error {
	// check password length
	if len(password) == 0 || len(password) < 6 {
		return errors.New("Password length must be greater than 5")
	}
	return nil
}

func ValidateUniqueEmail(email string) error {
	var user User
	db.Where("email = ?", email).First(&user)
	if user.Email == email {
		return errors.New("A user with this email exists.")
	}
	return nil
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
	}
	return db
}

func (user *User) CreateUser() (*User, error) {
	if mailErr := ValidateEmail(user.Email); mailErr != nil {
		return nil, mailErr
	}
	if passErr := ValidatePassword(user.Password); passErr != nil {
		return nil, passErr
	}
	// validate unique email
	if uniqueMail := ValidateUniqueEmail(user.Email); uniqueMail != nil {
		return nil, uniqueMail
	}
	// encrypt password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, hashErr
	}
	user.Password = string(hashedPassword)
	// create user
	user.ID = uuid.Must(uuid.NewV4()).String()
	db.Create(&user)
	return user, nil
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
		return nil, errors.New("Update unsuccessful. Bucketlist not found.")
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

// user

func GetUser(id string) (*User, error) {
	var user User
	db.Where("id = ?", id).First(&user)
	if user.ID == id {
		return &user, nil
	}
	return nil, errors.New("User not found.")
}

func Login(email, password string) (*Token, error) {
	// validate credentials
	var user User
	db.Where("email = ?", email).First(&user)
	if user.Email != email {
		return nil, errors.New("Email not found.")
	}
	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if hashErr != nil && hashErr == bcrypt.ErrMismatchedHashAndPassword {
		return nil, hashErr
	}
	// create JWT token
	tk := &Token{UserId: user.ID, Email: user.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, tkErr := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if tkErr != nil {
		return nil, tkErr
	}
	tk.AuthToken = tokenStr
	return tk, nil
}
