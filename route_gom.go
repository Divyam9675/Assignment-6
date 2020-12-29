package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {


	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=admin sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUser(w http.ResponseWriter, r *http.Request) {


	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=admin sslmode=disable")

	if err != nil {
		panic("could not find")
	}
	defer db.Close()

	var user []User

	db.Find(&user)
	json.NewEncoder(w).Encode(&user)

}

func NewUser(w http.ResponseWriter, r *http.Request) {


	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=admin sslmode=disable")

	if err != nil {
		panic("could not find")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	var user []User

	db.Find(&user)
	json.NewEncoder(w).Encode(&user)



}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=admin sslmode=disable")

	if err != nil {
		panic("could not find")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User

	db.Where("name =?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "user successfully deleted")

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=admin sslmode=disable")

	if err != nil {
		panic("could not find")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User

	db.Where("name =?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "successfully updated user")
}
