package main

import (
	"log"

	"github.com/wincentrtz/fake-news/config"
)

func createPostTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE posts(
			id serial PRIMARY KEY,
			post_title VARCHAR UNIQUE NOT NULL,
			post_description VARCHAR NOT NULL,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createPostStatusTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE post_status(
			id serial PRIMARY KEY,
			post_id integer NOT NULL,
			status integer NOT NULL,
			created_on TIMESTAMP NOT NULL,
			FOREIGN KEY (post_id) REFERENCES posts (id)
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE users(
			id serial PRIMARY KEY,
			name VARCHAR NOT NULL,
			email VARCHAR NOT NULL,
			password VARCHAR NOT NULL,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createPostTable()
	createPostStatusTable()
	createUserTable()
}
