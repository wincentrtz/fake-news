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

	postStatusHandler "github.com/wincentrtz/fake-news/domain/poststatus/handler/rest"
	_postStatusRepository "github.com/wincentrtz/fake-news/domain/poststatus/repository"
	_postStatusUsecase "github.com/wincentrtz/fake-news/domain/poststatus/usecase"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	r := mux.NewRouter()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	pr := _repository.NewPostRepository(db)
	pu := _usecase.NewPostUsecase(pr, timeoutContext)
	handler.NewPostHandler(r, pu)

	pqr := _postStatusRepository.NewPostStatusRepository(db)
	pqu := _postStatusUsecase.NewPostStatusUsecase(pqr, timeoutContext)
	postStatusHandler.NewPostStatusHandler(r, pqu)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
