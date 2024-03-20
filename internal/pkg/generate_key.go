package pkg

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/google/uuid"
)

// GenerateAccessKey 生成 用户的唯一标识（访问密钥）
func GenerateAccessKey() string {
	return uuid.NewString()
}

// GenerateSecureSecretKey 生成一个指定长度的安全的secretKey
func GenerateSecureKey() (string, error) {
	var byteLength int = 32
	// 创建一个足够大的空间来保存随机字节
	randomBytes := make([]byte, byteLength)
	// 使用crypto/rand生成随机字节
	if _, err := rand.Read(randomBytes); err != nil {
		// 如果生成过程中发生错误，返回错误
		return "", err
	}
	// 将随机字节序列编码为十六进制字符串
	secretKey := hex.EncodeToString(randomBytes)
	return secretKey, nil
}
