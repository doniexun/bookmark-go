package utils

import (
	"log"
	"github.com/mojocn/base64Captcha"
)

// 生成
func GetCaptcha() (str1 string, str2 string) {
	var configD = base64Captcha.ConfigDigit{
		Height:     40,
		Width:      120,
		MaxSkew:    0.9,
		DotCount:   80,
		CaptchaLen: 5,
	}

	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	return idKeyD, base64stringD
}

// 校验（验证成功后自动失效）
func VerfiyCaptcha(idkey string, verifyValue string) bool {
	log.Println("验证码校验", idkey, verifyValue)

	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
    if verifyResult {
        return true
    } else {
        return false
    }
}
