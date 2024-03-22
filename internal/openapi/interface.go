package openapi

import (
	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/internal/common/openrequest"
	"github.com/showmebug/my-gin-demo/internal/common/response"
)

type WeatherData struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Date    string `json:"date"`
	Type    string `json:"type"`
	Weather string `json:"weather"`
}

// 查询天气 （模拟接口）
func QueryWeather(c *gin.Context) {
	var form openrequest.Weather
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, openrequest.GetErrorMsg(form, err))
		return
	}
	data := WeatherData{
		Country: "China",
		City:    "Beijing",
		Date:    "2024-03-22",
		Type:    "Current",
		Weather: "Sunny",
	}
	response.Success(c, data)
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	response.BusinessFail(c, err.Error())
	// } else {
	// 	response.Success(c, jsonData)
	// }
}
