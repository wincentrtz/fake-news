package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/wincentrtz/fake-news/config"
	handler "github.com/wincentrtz/fake-news/domain/post/handler/rest"
	_repository "github.com/wincentrtz/fake-news/domain/post/repository"
	_usecase "github.com/wincentrtz/fake-news/domain/post/usecase"

	postQueueHandler "github.com/wincentrtz/fake-news/domain/postqueue/handler/rest"
	_postQueueRepository "github.com/wincentrtz/fake-news/domain/postqueue/repository"
	_postQueueUsecase "github.com/wincentrtz/fake-news/domain/postqueue/usecase"
	"github.com/wincentrtz/fake-news/models/builder"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	post := builder.NewPost().Title("Title 1").Author("David").Content("asdasdadsadas").Build()
	r := mux.NewRouter()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	fmt.Println(post)

	pr := _repository.NewPostRepository(db)
	pu := _usecase.NewPostUsecase(pr, timeoutContext)

	pqr := _postQueueRepository.NewPostQueueRepository(db)
	pqu := _postQueueUsecase.NewPostQueueUsecase(pqr, timeoutContext)

	handler.NewPostHandler(r, pu)
	postQueueHandler.NewPostQueueHandler(r, pqu)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
