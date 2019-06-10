package main

import (
	"log"

	"github.com/wincentrtz/fake-news/config"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	schema := "DROP TABLE posts, postqueues"
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}
