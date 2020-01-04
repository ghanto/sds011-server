package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ghanto/sds011-server/pkg/sensor/infrastructure"

	"github.com/caarlos0/env"

	"github.com/ghanto/sds011-server/config"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	appConfig := config.ApplicationConfig

	if err := env.Parse(&appConfig); err != nil {
		fmt.Printf("%+v\n", err)
	}

	setupDb(appConfig)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

	fmt.Println("Server started on 9999")
}

func setupDb(appConfig config.Config) {
	db, err := infrastructure.NewPostgresDb(appConfig.DatabaseDSN)
	if err != nil {
		panic(err)
	}
	if err := db.RunMigrations(); err != nil {
		panic(err)
	}
}
