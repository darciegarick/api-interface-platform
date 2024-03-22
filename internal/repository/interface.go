package repository

import "strconv"

type Interface struct {
	ID
	Name           string `json:"name" gorm:"not null;comment:接口名称"`
	Description    string `json:"description" gorm:"comment:接口描述"`
	Url            string `json:"url" gorm:"not null;comment:接口地址"`
	RequestHeader  string `json:"request_header" gorm:"comment:请求头"`
	ResponseHeader string `json:"response_header" gorm:"comment:响应头"`
	Status         int    `json:"status" gorm:"default:1;comment:接口状态"`
	MethodType     string `json:"method_type" gorm:"default:GET;comment:请求类型"`
	UserId         int    `json:"user_id" gorm:"comment:用户ID"`
	Timestamps
	SoftDeletes
}

func (itface Interface) GetUid() string {
	return strconv.Itoa(int(itface.ID.ID))
}
