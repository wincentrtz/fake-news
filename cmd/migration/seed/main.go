package main

import (
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/wincentrtz/fake-news/config"
	"github.com/wincentrtz/fake-news/models"
)

func insertUserDataToDB(user models.User) {
	db := config.InitDb()
	defer db.Close()

	query := `INSERT INTO users (name, email, password, created_on)
		VALUES($1,$2,$3,$4)
	`
	_, err := db.Exec(query, &user.Name, &user.Email, &user.Password, time.Now())

	if err != nil {
		panic(err)
	}
}

func populateUserData(number int) {
	for i := 0; i < number; i++ {
		user := models.User{}
		err := faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
		}
		insertUserDataToDB(user)
	}
}

func main() {
	populateUserData(100)
}
