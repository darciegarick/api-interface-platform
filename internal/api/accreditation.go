package api

import (
	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/internal/common/request"
	"github.com/showmebug/my-gin-demo/internal/common/response"
)

// GetSign 获取接口 的验证签名
func GetSign(c *gin.Context) {
	var form request.GetSign
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

}
