package main

import (
	"fmt"

	"github.com/wincentrtz/fake-news/post/models/builder"
)

func main() {
	post := builder.NewPost().Title("Title 1").Author("David").Content("asdasdadsadas").Build()
	fmt.Println(post)
}
