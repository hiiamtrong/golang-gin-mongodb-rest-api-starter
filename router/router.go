package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/middleware"
	v1 "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router/api/v1"
)

func InitRouter(
	config *config.Config,
	todoAPIV1 *v1.TodoAPIV1,
	authAPIV1 *v1.AuthAPIV1,
) *gin.Engine {
	gin.SetMode(config.App.Env.String())
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	m := middleware.NewAppMiddleware(config)
	SetupSwagger(r, config)

	pingAPI := r.Group("/ping")
	{
		pingAPI.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	v1 := r.Group("/api/v1")
	{

		authV1 := v1.Group("/auth")
		{
			authV1.POST("/register", authAPIV1.Register())
			authV1.POST("/login", authAPIV1.Login())
		}

		todoV1 := v1.Group("/todo")
		{
			todoV1.Use(m.JWT())
			todoV1.GET("", todoAPIV1.List())
			todoV1.POST("", todoAPIV1.Create())
			todoV1.PUT("/:id", todoAPIV1.Update())
			todoV1.DELETE("/:id", todoAPIV1.Delete())
		}
	}

	return r
}
