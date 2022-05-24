package repository

import (
	"database/sql"
	"log"

	"github.com/basic/model"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func CreateConnection() {
	log.Print("..Creating database...")
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("cannot open an SQLite memory database: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE USER (
		UID INTEGER PRIMARY KEY AUTOINCREMENT,
		USERNAME VARCHAR(64) NULL, 
		CREATED DATE NULL);`)
	if err != nil {
		log.Fatalf("cannot create schema: %v", err)
	}

	DB = db
	log.Println("Success!!")
}

func CreateUser(newUser *model.User) (*int64, error) {
	log.Printf("..Creating User %v...", newUser)
	res, err := DB.Exec("INSERT INTO USER(username, created) values(?,?); ", newUser.Username, newUser.Created)
	if err != nil {
		log.Fatalf("cannot insert data: %v", err)
		return nil, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return nil, err
	}
	return &id, nil
}

func GetAllUsers() ([]model.User, error) {
	rows, err := DB.Query("SELECT * FROM USER; ")
	if err != nil {
		return nil, err
	}

	returnUsers := []model.User{}

	for rows.Next() {
		userParsed := model.User{}
		err = rows.Scan(&userParsed.UID, &userParsed.Username, &userParsed.Created)
		if err != nil {
			return nil, err
		}

		returnUsers = append(returnUsers, userParsed)
	}
	return returnUsers, nil
}

func GetUserByID(id int) (*model.User, error) {
	userFound := model.User{}
	row := DB.QueryRow("SELECT UID, USERNAME, CREATED FROM USER WHERE UID=?", id)
	if err := row.Scan(&userFound.UID, &userFound.Username, &userFound.Created); err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	return &userFound, nil
}
