package main

import (
	"fmt"

	"github.com/wincentrtz/fake-news/models/builder"
)

func main() {
	user := builder.NewUser().Name("David").Build()
	fmt.Println(user.Name)
}
