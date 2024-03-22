package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/internal/openapi"
)

func SetOpenApiGroupRoutes(router *gin.RouterGroup) {

	router.POST("/getWeather", openapi.QueryWeather)

}
