package middleware

import "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"

type AppMiddleware struct {
	config *config.Config
}

func NewAppMiddleware(config *config.Config) *AppMiddleware {
	return &AppMiddleware{
		config: config,
	}
}
