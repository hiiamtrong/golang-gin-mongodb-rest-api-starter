package router

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/docs/swagger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func SetupSwagger(r *gin.Engine, config *config.Config) {
	swgCfg := config.Swagger
	swagger.SwaggerInfo.Title = "Todo API"
	swagger.SwaggerInfo.Description = "This is a sample server Todo server."
	swagger.SwaggerInfo.Version = "1.0"
	swagger.SwaggerInfo.Host = swgCfg.Host
	swagger.SwaggerInfo.BasePath = swgCfg.BasePath
	swagger.SwaggerInfo.Schemes = strings.Split(swgCfg.Schemes, ",")

	group := r.Group("/swagger", gin.BasicAuth(gin.Accounts{
		swgCfg.Username: swgCfg.Password,
	}))

	group.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
