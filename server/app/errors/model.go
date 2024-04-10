package errors

type ErrorCode int

// 常规码
const (
	CodeSuccess      ErrorCode = 200
	CodeApiAuthFail  ErrorCode = 400
	CodeTokenInvalid ErrorCode = 401
	CodeServerError  ErrorCode = 500
)

// 业务码
const (
	ErrorCodeParamsError              ErrorCode = 100100 + iota // 参数有误
	ErrorCodeCaptchaInvalid                                     // 验证码无效
	ErrorCodeAccountOrPasswordInvalid                           // 账号或密码不正确
	ErrorCodeAccountNotActivate                                 // 账号未激活
	ErrorCodePasswordsNotMatch                                  // 密码不一致
	ErrorCodeOldPasswordInvalid                                 // 旧密码错误
	ErrorCodeHasNotPermission                                   // 没有权限
	ErrorCodeAccountAlreadyExist                                // 账号已存在
	ErrorCodeGameAlreadyExist                                   // 游戏已存在
	ErrorCodeCanNotChangeSelf                                   // 不可修改自身账号
	ErrorCodeFileExtNotMacth                                    // 文件格式不符
)

var (
	// 错误码对应信息
	Msgs = map[ErrorCode]string{
		CodeSuccess:                       "success",
		CodeApiAuthFail:                   "api auth fail",
		CodeTokenInvalid:                  "token invalid",
		CodeServerError:                   "server error",
		ErrorCodeParamsError:              "params error",
		ErrorCodeCaptchaInvalid:           "captcha invalid",
		ErrorCodeAccountOrPasswordInvalid: "account or password invalid",
		ErrorCodeAccountNotActivate:       "account not activate",
		ErrorCodePasswordsNotMatch:        "new password and confirm password not match",
		ErrorCodeOldPasswordInvalid:       "old password invalid",
		ErrorCodeHasNotPermission:         "has not permission",
		ErrorCodeAccountAlreadyExist:      "account already exist",
		ErrorCodeGameAlreadyExist:         "game already exist",
		ErrorCodeCanNotChangeSelf:         "can not change self account",
		ErrorCodeFileExtNotMacth:          "file ext not macth",
	}
)
