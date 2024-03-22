package services

import (
	"github.com/showmebug/my-gin-demo/internal/common/request"
)

type accreditationService struct {
}

var AccreditationService = new(accreditationService)

// createSign 生成签名
func (accreditationService *accreditationService) createSign(params request.GetSign) (err error, sign string) {
	//
	return
}

// func (accreditationService *accreditationService) createSign(accessKey, secretKey, nonce string, timestamp int64) string {
// 	signStr := pkg.GenerateSign(accessKey, secretKey, nonce, timestamp)
// 	return signStr
// }

// verifySign 校验签名
// func (accreditationService *accreditationService) verifySign(signature, accessKey) bool {

// }
