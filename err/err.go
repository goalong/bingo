package err

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	PostNotFound = 1001

	UserNotFound = 1101
	EMAIL_OR_PW_WRONG = 1102

	SignTokenError = 1201
	CHECK_TOKEN_FAILED = 1202
	CHECK_TOKEN_TIMEOUT = 1203

	CONTEXT_GET_USER_ERROR = 1301
	PERMISSION_DENIED = 1401


)

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	PostNotFound: "文章不存在",
	UserNotFound: "用户不存在",
	SignTokenError: "签发token错误",
	CHECK_TOKEN_FAILED: "验证token时失败",
	CHECK_TOKEN_TIMEOUT: "token已过期",
	EMAIL_OR_PW_WRONG: "邮箱不存在或者密码错误",
	CONTEXT_GET_USER_ERROR: "上下文错误",
	PERMISSION_DENIED: "权限不足",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}




