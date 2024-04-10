package captcha

import "github.com/mojocn/base64Captcha"

// 生成验证码图片
func CreateCode() (string, string, string, error) {
	return base64Captcha.NewCaptcha(captchaConfig, captchaResult).Generate()
}

// 校验验证码
func VerifyCaptcha(code string, captcha string) bool {
	return captchaResult.Verify(code, captcha, true)
}
