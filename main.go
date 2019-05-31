package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/fake-news/post/handlers"
	"github.com/wincentrtz/fake-news/post/models/builder"
)

func main() {
	// config.InitDb()
	post := builder.NewPost().Title("Title 1").Author("David").Content("asdasdadsadas").Build()
	r := mux.NewRouter()
	fmt.Println(post)

	handlers.NewPostHandler(r)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
