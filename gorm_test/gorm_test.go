package gorm_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

type Dog struct {
	// customized the Primary Key, the `type` tag is the data type in the table
	UUID uuid.UUID `gorm:"primaryKey;autoIncrement:false;type:char(36)"`
	Name string
	Toys []Toy `gorm:"foreignKey:DogUUID"`
	gorm.DeletedAt
}

type Toy struct {
	gorm.Model
	Name    string
	DogUUID string `gorm:"references:UUID"`
}

// BeforeCreate is the gorm.DB hook before insert into database
func (dog *Dog) BeforeCreate(tx *gorm.DB) (err error) {
	dog.UUID = uuid.New()
	return nil
}

var fapp = faker.New().App()

// GenerateDogs func used to create faker data.
func GenerateDogs(n int) []Dog {
	var dogs []Dog
	var dog Dog
	for i := 0; i < n; i++ {
		dog = Dog{}
		dog.Name = fapp.Faker.Pet().Dog()
		dogs = append(dogs, dog)
	}
	return dogs
}

// SaveToDB uses variadic parameters for one or many dogs.
func SaveToDB(db *gorm.DB, dogs ...Dog) (err error) {
	db.Create(&dogs)
	return nil
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

func Test_Has_Many(t *testing.T) {
	AutoMigrate()
	db, err := MakeDbConnection()
	if err != nil {
		t.Fatal("Can not connect to the database.")
	}

	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	// SaveToDB func use variadic parameters , and in batch mode insert into DB
	// SaveToDB(db, GenerateDogs(10)...)
	// SaveToDB(db, dog1, dog2, ... dogN)

	/*************************************************************
	*          One-To-Many Query in MySQL
	**************************************************************/

	// Inner join two tables may has the name of fields are conflicted,
	// you can use the `as` keyword give the field an alias.

	var count int64

	rows, err := db.Model(&Dog{}).Select("dogs.uuid, dogs.name as Dog, toys.name as Toy").
		Joins("inner join toys ON uuid = dog_uuid").
		// SQL uses single `=`
		Where("dogs.uuid = ? ", "5247e371-d046-4478-baa4-f1d1cb249803").
		Count(&count).
		Rows()
	if err != nil {
		t.Log("failed to fetch data from joined tables")
	}

	// The Result struct must have the same name of attributes  according to the results.
	// otherwise the db.ScanRows() func will ignore the attribute which is not matched any
	// of the fileds name in the query result.

	type Result struct {
		Dog  string
		UUID string
		Toy  string
	}

	// var result Result
	for rows.Next() {
		var row Result
		db.ScanRows(rows, &row)
		t.Logf(`%s [%s] has %d toys "%s"`, row.Dog, row.UUID, count, row.Toy)
	}
}

func Test_UUID(t *testing.T) {
	t.Log(len(uuid.New().String()))
}

func Test_Faker(t *testing.T) {
	fapp := faker.New().App()
	for i := 0; i < 20; i++ {
		t.Log(fapp.Faker.Person().FirstName())
	}
}
