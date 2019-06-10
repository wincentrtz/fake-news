package main

import (
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
)

func main() {
	db := config.InitDb()
	defer db.Close()
	r := mux.NewRouter()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	pr := _repository.NewPostRepository(db)
	pu := _usecase.NewPostUsecase(pr, timeoutContext)

	pqr := _postQueueRepository.NewPostQueueRepository(db)
	pqu := _postQueueUsecase.NewPostQueueUsecase(pqr, timeoutContext)

	handler.NewPostHandler(r, pu)
	postQueueHandler.NewPostQueueHandler(r, pqu)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}