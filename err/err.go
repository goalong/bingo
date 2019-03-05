package err

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	PostNotFound = 1001

)

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	PostNotFound: "文章不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}




