package main

import (
	"context"
	"fmt"
	httpDelivery "github.com/asdamwongmantap/api-gomongo-crud/crud/delivery/http"
	"github.com/asdamwongmantap/api-gomongo-crud/crud/model"
	"github.com/asdamwongmantap/api-gomongo-crud/crud/repository"
	"github.com/asdamwongmantap/api-gomongo-crud/crud/usecase"
	"github.com/asdamwongmantap/api-gomongo-crud/lib/config"
	"github.com/asdamwongmantap/api-gomongo-crud/lib/db"
	mongoDB "github.com/asdamwongmantap/api-gomongo-crud/lib/db/mongo"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	config.SetConfigFile("config", "config", "json")
}

func main() {
	envConfig := getConfig()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Mongo
	mongo, err := mongoDB.NewV2(context.Background(), envConfig.Mongo)
	if err != nil {
		log.Println(err)
		return
	}

	crudRepo := repository.NewCrudRepository(mongo)
	crudUseCase := usecase.NewCrudUseCase(&envConfig, crudRepo)
	// Router
	httpDelivery.NewRouter(e, crudUseCase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.GetString("host.address"), config.GetString("host.port"))))
}

func getConfig() model.EnvConfig {

	return model.EnvConfig{
		AppName: config.GetString("name"),
		Env:     config.GetString("env"),

		Mongo: db.MongoConfig{
			Timeout:  config.GetInt("database.mongodb.timeout"),
			DBname:   config.GetString("database.mongodb.dbname"),
			Username: config.GetString("database.mongodb.user"),
			Password: config.GetString("database.mongodb.password"),
			Host:     config.GetString("database.mongodb.host"),
			Port:     config.GetString("database.mongodb.port"),
		},
	}
}
