package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
)

//Db is imported from the sql package creates and frees connections automatically
var Db *sql.DB

//InitDB Function creates a connection to our database
func InitDB() {
	db, err := sql.Open("mysql", "root:Newyork@123@tcp(127.0.0.1:3306)/dogsdb")
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

//Migrate function runs migrations file for us
func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
