package api

import (
	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/internal/common/request"
	"github.com/showmebug/my-gin-demo/internal/common/response"
	"github.com/showmebug/my-gin-demo/internal/services"
)

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}
