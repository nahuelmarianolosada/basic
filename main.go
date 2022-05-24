package main

import (
	"log"
	"time"

	"github.com/basic/model"
	"github.com/basic/repository"
)

var usersToInsert = []model.User{
	{Username: "Elthon John", Created: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
	{Username: "Jimmy Page", Created: time.Date(2010, 06, 25, 22, 01, 00, 0, time.UTC)},
	{Username: "Roberto Sanchez", Created: time.Now()},
}

func main() {
	repository.CreateConnection()

	users, err := repository.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)

	for _, user := range usersToInsert {
		idCreated, err := repository.CreateUser(&user)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("ID del usuario creado: %d \n", *idCreated)
	}
	users, err = repository.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		log.Printf("User getted %v\n", user)
	}

	if userFound, err := repository.GetUserByID(2); err != nil {
		log.Println(err)
	} else {
		log.Printf("User Finded: %v", userFound)
	}
}
