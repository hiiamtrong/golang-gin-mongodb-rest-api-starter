package main

import (
	"log"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/database"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/repository"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router"
	v1 "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router/api/v1"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/server"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/service"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config.NewConfig: %v", err)
	}

	mongo := database.NewMongodb(config)

	todoRepo := repository.NewTodoRepository(mongo)
	authRepo := repository.NewAuthRepository(mongo)

	todoSvc := service.NewTodoService(todoRepo)
	authSvc := service.NewAuthService(config, authRepo)

	todoAPIV1 := v1.NewTodoAPIV1(todoSvc)
	authAPIV1 := v1.NewAuthAPIV1(authSvc)

	router := router.InitRouter(
		config,
		todoAPIV1,
		authAPIV1,
	)

	httpServer := server.NewServer(config, router)

	if err := httpServer.Run(); err != nil {
		log.Fatalf("server.Run: %v", err)
	}
}
