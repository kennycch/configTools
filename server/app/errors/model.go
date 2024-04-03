package errors

type ErrorCode int

// 常规码
const (
	CodeSuccess     ErrorCode = 200
	CodeClientError ErrorCode = 400
	CodeServerError ErrorCode = 500
)

// 业务码
const (
	ErrorCodeApiAuthFail            ErrorCode = 100100 + iota // 前端接口签名不通过
	ErrorCodeParamsError                                   // 参数有误
	ERROR_CODE_TOKEN_IS_INVALID                               // 令牌无效
	ERROR_CODE_PLEASE_LOGIN                                   // 没登录
	ERROR_CODE_DUPLICATE_REQUEST                              // 并发请求
	ERROR_CODE_NO_TIMES_CAN_RECEIVE                           // 没有可领取次数
	ERROR_CODE_HAS_NOT_PLAY_TIMES                             // 没有挑战/复活次数
	ERROR_CODE_IN_BLACK_LIST                                  // 在黑名单中
	ERROR_CODE_ALREADY_RESERVATION                            // 已预约
	ERROR_CODE_ROLE_ALREADY_BINDING                           // 角色已绑定
	ERROR_CODE_NO_GIFT_CAN_RECEIVE                            // 没有礼包可领取
	ERROR_CODE_SERVER_NOT_ENABLE                              // 区组不可用
	ERROR_CODE_ACTIVITY_NOT_START                             // 活动没开始
	ERROR_CODE_ACTIVITY_ALREADY_END                           // 活动已结束
	ERROR_CODE_INVALID_REPORT                                 // 无效上报
	ERROR_CODE_CAN_NOT_ACCESS                                 // 不可访问
	ERROR_CODE_BLACK_INDUSTRY                                 // 黑产账号
)

var (
	// 错误码对应信息
	Msgs = map[ErrorCode]string{
		CodeSuccess:     "success",
		CodeClientError: "client error",
		CodeServerError: "server error",
	}
)
