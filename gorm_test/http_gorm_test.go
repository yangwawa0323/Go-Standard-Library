package gorm_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/mux"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string `gorm:"name"`
	Email string `gorm:"email"`
}

func MakeDbConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open("root:redhat@tcp(127.0.0.1:3306)/go_standard_library?charset=utf8mb4&parseTime=True&loc=Local"))
	return
}

func InitialMigration() {
	db, err := MakeDbConnection()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Failed to connect to database")
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	db.AutoMigrate(&User{})
}

// Handler func
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = MakeDbConnection()
	if err != nil {
		panic("Could not connect to the database")
	}

	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)

}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err := MakeDbConnection()
	if err != nil {
		panic("Could not connect to the database")
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	newUser := &User{Name: name, Email: email}

	db.Create(newUser)

	fmt.Fprintf(w, "New User %s with id: %d Successfully created .", newUser.Name, newUser.ID)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := MakeDbConnection()
	if err != nil {
		panic("Could not connect to the database")
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	db.Where("id = ?", id).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "Delete User %s Successfully deleted.", user.Name)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := MakeDbConnection()
	if err != nil {
		panic("Could not connect to the database")
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	// newUser := &User{Name: name, Email: email}

	var user User

	db.Where("name = ?", name).First(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "Update User Successfully updated")
}

//////////////////////////////

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func Test_Mux_Router(t *testing.T) {
	t.Log("go-mux demo")
	InitialMigration()
	handleRequest()
}
