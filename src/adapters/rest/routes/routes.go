package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/chucknorris/src/adapters/rest/controllers"
	"github.com/spf13/viper"
)

//SetupRouter sets up routes
func SetupRouter() *gin.Engine {
	gin.SetMode(viper.GetString("server.mode"))
	router := gin.Default()
	api := router.Group("/api/v1.0")
	{
		api.GET("/ping", controllers.Pong)
		api.GET("/call", controllers.Call)
	}
	return router
}
