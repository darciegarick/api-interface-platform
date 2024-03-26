package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/internal/api"
	"github.com/showmebug/my-gin-demo/internal/middleware"
	"github.com/showmebug/my-gin-demo/internal/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	router.POST("/auth/register", api.Register)

	router.POST("/auth/login", api.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", api.Info)
		// authRouter.POST("/auth/logout", api.Logout)
	}

	// =======================================

}
