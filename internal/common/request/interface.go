package request

type GetSign struct {
	Interface  string `form:"interface" json:"interface" binding:"required"`
	AccessKey  string `form:"access_key" json:"access_key" binding:"required"`
	Identifier string `form:"identifier" json:"identifier" binding:"required"`
	Timestamp  string `form:"timestamp" json:"timestamp" binding:"required"`
}

func (getSign GetSign) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"interface.required":  "接口不能为空",
		"access_key.required": "AccessKey不能为空",
		"identifier.required": "标识不能为空",
		"timestamp.required":  "时间戳不能为空",
	}
}
