package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"
	timepkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/time"
)

type Server struct {
	HTTPServer *http.Server
}

func NewServer(
	config *config.Config,
	router *gin.Engine,
) *Server {
	address := fmt.Sprintf(":%s", config.Server.Port)
	readTimeout := timepkg.StrToDuration(config.Server.ReadTimeout)
	writeTimeout := timepkg.StrToDuration(config.Server.WriteTimeout)

	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return &Server{
		HTTPServer: server,
	}
}

func (s *Server) Run() error {
	log.Println("Server is running on", s.HTTPServer.Addr)
	return s.HTTPServer.ListenAndServe()
}
