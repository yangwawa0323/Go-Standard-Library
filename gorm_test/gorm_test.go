package gorm_test

import (
	"fmt"
	"log"
	"testing"
)

type Dog struct {
	ID   int
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

func AutoMigrate() {
	db, err := MakeDbConnection()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Failed to connect to database")
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	db.AutoMigrate(&Dog{}, &Toy{})
}

func Test_Has_Many_Polymorphism(t *testing.T) {
	AutoMigrate()
	db, err := MakeDbConnection()
	if err != nil {
		t.Fatal("Can not connect to the database.")
	}

	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	// db.Create(&Dog{Name: "dog1", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})

	var dog Dog
	db.Model(&Dog{}).Preload("Toys").First(&dog)
	t.Logf("Dog %v", dog)
}
