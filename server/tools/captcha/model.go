package captcha

import (
	"image/color"
	"time"

	"github.com/mojocn/base64Captcha"
)

var (
	// 验证码配置
	captchaConfig = &base64Captcha.DriverString{
		Height:          100,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          6,
		Source:          "0123456789",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	// 验证码图片存储规则
	captchaResult = base64Captcha.NewMemoryStore(20480, 5*time.Minute)
)
