package dogs

import (
	"fmt"
	"log"

	database "github.com/alekhyakamale/go-api/internal/pkg/db/migrations/mysql"
)

//Dog struct defined to represent the databse schema
type Dog struct {
	ID        string
	Name      string
	IsGoodBoi bool
}

//Save our dog object in database and returns ID
func (dog Dog) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Dogs(Name,IsGoodBoi) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(dog.Name, dog.IsGoodBoi)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

//When the PREPARE statement is executed, the specified statement is parsed,
//analyzed, and rewritten. When an EXECUTE command is subsequently issued,
//the prepared statement is planned and executed. This division of labor avoids
//repetitive parse analysis work, while allowing the execution plan to depend on
//the specific parameter values supplied.

//GetAll our dogs
func GetAll() []Dog {
	stmt, err := database.Db.Prepare("select id, name, isGoodBoi from dogs")
	if err != nil {
		log.Fatal(err)
	}

	//Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	//Close returns the connection to the connection pool.
	defer rows.Close()
	var dogs []Dog
	for rows.Next() {
		var dog Dog
		err := rows.Scan(&dog.ID, &dog.Name, &dog.IsGoodBoi)
		if err != nil {
			log.Fatal(err)
		}
		dogs = append(dogs, dog)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return dogs
}

//UpgradeDog changes the status of the dog
func (dog Dog) UpgradeDog() {
	stmt, err := database.Db.Prepare("UPDATE Dogs SET isGoodBoi=? WHERE name=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(dog.IsGoodBoi, dog.Name)
	if err != nil {
		log.Fatal(err)
	}
	result, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	fmt.Println(result)
	log.Print(dog.Name + " was upgraded")
}

//UpForAdoption deletes the dogs from db
func (dog Dog) UpForAdoption() {
	stmt, err := database.Db.Prepare("DELETE FROM Dogs WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(dog.ID)
	if err != nil {
		log.Fatal(err)
	}
	result, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	fmt.Println(result)
}
