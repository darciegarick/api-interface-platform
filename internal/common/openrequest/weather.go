package openrequest

type Weather struct {
	Country string `form:"country" json:"country" binding:"required"`
	City    string `form:"city" json:"city" binding:"required"`
	Date    string `form:"date" json:"date" binding:"required"`
	Tyoe    string `form:"type" json:"type" binding:"required"`
}

func (weather Weather) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"country.required": "国家名称不能为空",
		"city.required":    "城市名称不能为空",
		"date.required":    "查询日期不能为空",
		"type.required":    "查询类型不能为空",
	}
}
