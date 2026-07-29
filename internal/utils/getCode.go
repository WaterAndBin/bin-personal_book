package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// 生成6位数字验证码
func GenerateCode() (string, error) {
	code := ""

	for i := 0; i < 6; i++ {

		// 生成 0-9 之间的安全随机整数
		// rand.Reader 密码学安全的随机数源
		// 表示随机数的上限（即生成 0 到 9 之间的数）
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}

		code += strconv.FormatInt(num.Int64(), 10)
	}

	return code, nil
}
