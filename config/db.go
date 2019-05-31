package config

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/lib/pq" //pq driver _ import
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func InitDb() {
	host := viper.GetString(`database.host`)
	port := viper.GetString(`database.port`)
	user := viper.GetString(`database.user`)
	password := viper.GetString(`database.password`)
	dbname := viper.GetString(`database.dbname`)
	fmt.Printf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
