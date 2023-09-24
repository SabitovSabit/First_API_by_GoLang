package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type User struct {
	Id   int
	Name string
}



func Main() {
	server := "DESKTOP-LG63R1M"
	port := 1433
	database := "SchoolDb"

	connString := fmt.Sprintf("server=%s;port=%d;database=%s;integrated security=true;", server, port, database)

	db, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected!")

	InsertUser(db,err)
}

func GetAllUsers(db *sql.DB,err error) []User {

	rows, err := db.Query("SELECT id, Name FROM Teachers")
	fmt.Println("Teachers selected!")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	fmt.Println("Users:", users)

	return users
}

func InsertUser(db *sql.DB,err error){
	query:=fmt.Sprint("insert into Teachers (name) values(?)")

	stmt,err:=db.Prepare(query)
    if err != nil {
		log.Fatal(err)
	} 
	defer stmt.Close()

	_,err=stmt.Exec("TestGo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
}
