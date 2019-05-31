package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/wincentrtz/fake-news/config"
	handler "github.com/wincentrtz/fake-news/domain/post/handler/rest"
	_postRepository "github.com/wincentrtz/fake-news/domain/post/repository"
	_postUsecase "github.com/wincentrtz/fake-news/domain/post/usecase"
	"github.com/wincentrtz/fake-news/models/builder"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	post := builder.NewPost().Title("Title 1").Author("David").Content("asdasdadsadas").Build()
	r := mux.NewRouter()
	fmt.Println(post)

	pr := _postRepository.NewPostRepository(db)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	pu := _postUsecase.NewPostUsecase(pr, timeoutContext)

	handler.NewPostHandler(r, pu)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
