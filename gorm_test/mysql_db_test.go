package gorm_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jaswdr/faker"
)

type ConnURI struct {
	User            string
	Secret          string
	Host            string
	Port            string
	DefaultDatabase string // Go struct cannot has default value neither in func.
}

type Student struct {
	StudentId int
	Name      string
	Age       int
}

// MySQLConn struct is proxy of MysQL database connection
type MySQLConn struct {
	Db *sql.DB
}

func (conn *MySQLConn) Exec(query string) (result sql.Result, err error) {
	return conn.Db.Exec(query)
}

func (conn *MySQLConn) Query(query string) (rows *sql.Rows, err error) {
	return conn.Db.Query(query)
}

func (conn *MySQLConn) Close() error {
	return conn.Db.Close()
}

func (conn *MySQLConn) MakeConn(uri ConnURI) *MySQLConn {

	if uri.Port == "" {
		uri.Port = "3306"
	}

	if uri.DefaultDatabase == "" {
		uri.DefaultDatabase = "test"
	}

	conn_uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		uri.User,
		uri.Secret,
		uri.Host,
		uri.Port,
		uri.DefaultDatabase,
	)

	db, err := sql.Open("mysql", conn_uri)
	if err != nil {
		log.Fatal("Can not make connection to database.")
	}
	conn.Db = db
	return conn
}

var connUri ConnURI = ConnURI{
	User:            "root",
	Secret:          "redhat",
	Port:            "3306",
	Host:            "127.0.0.1",
	DefaultDatabase: "go_standard_library",
}

func Test_MySQL_Db_Exec(t *testing.T) {
	t.Log(" MySQL Db Exec demo")

	var myConn *MySQLConn = new(MySQLConn)

	// MakeConn return a myConn pointer for chainning the func
	// myConn = myConn.MakeConn(connUri)
	myConn.MakeConn(connUri)

	defer myConn.Close()

	_, err := myConn.Exec(`CREATE TABLE IF NOT EXISTS student (
			sid INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
			name VARCHAR(30),
			age SMALLINT UNSIGNED
	)`)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Successfully connected to MySQL database.")

}

func Test_MySQL_Db_Query_Select(t *testing.T) {
	t.Log(" MySQL Db Query Demo")

	var myConn = new(MySQLConn).MakeConn(connUri)

	defer myConn.Close()

	query_str := `SELECT * FROM student LIMIT 10`
	rows, err := myConn.Query(query_str)
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var student Student = Student{}
		rows.Scan(&student.StudentId, &student.Name, &student.Age)
		t.Logf("ID: %d, Name: %s, Age: %d\n",
			student.StudentId, student.Name, student.Age)
	}
}

func Test_MySQL_Db_Query_Insert(t *testing.T) {
	t.Log(" MySQL Db Query Insert Demo")

	var myConn = new(MySQLConn).MakeConn(connUri)

	defer myConn.Close()

	faker := faker.New()

	for i := 1; i < 10; i++ {
		query_str := fmt.Sprintf(`INSERT INTO student(name,age) VALUES(
				'%s',%d
		)`,
			faker.Person().Name(),
			faker.IntBetween(19, 65),
		)
		_, err := myConn.Query(query_str)
		if err != nil {
			t.Fatal(err)
		}
	}

}
