package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/wincentrtz/fake-news/config"

	postHandler "github.com/wincentrtz/fake-news/domain/post/handler/rest"
	_postRepository "github.com/wincentrtz/fake-news/domain/post/repository"
	_postUsecase "github.com/wincentrtz/fake-news/domain/post/usecase"

	postStatusHandler "github.com/wincentrtz/fake-news/domain/poststatus/handler/rest"
	_postStatusRepository "github.com/wincentrtz/fake-news/domain/poststatus/repository"
	_postStatusUsecase "github.com/wincentrtz/fake-news/domain/poststatus/usecase"

	shareHandler "github.com/wincentrtz/fake-news/domain/share/handler/rest"
	_shareRepository "github.com/wincentrtz/fake-news/domain/share/repository"
	_shareUsecase "github.com/wincentrtz/fake-news/domain/share/usecase"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	r := mux.NewRouter()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	registerPostHandler(r, timeoutContext, db)
	registerStatusHandler(r, timeoutContext, db)
	registerShareHandler(r, timeoutContext, db)

	fmt.Println("Starting..")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func registerPostHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	pr := _postRepository.NewPostRepository(db)
	pu := _postUsecase.NewPostUsecase(pr, timeoutContext)
	postHandler.NewPostHandler(r, pu)
}

func registerStatusHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	pqr := _postStatusRepository.NewPostStatusRepository(db)
	pqu := _postStatusUsecase.NewPostStatusUsecase(pqr, timeoutContext)
	postStatusHandler.NewPostStatusHandler(r, pqu)
}

func registerShareHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	sr := _shareRepository.NewShareRepository(db)
	su := _shareUsecase.NewShareUsecase(sr, timeoutContext)
	shareHandler.NewShareHandler(r, su)
}
