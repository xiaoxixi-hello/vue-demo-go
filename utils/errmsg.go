package utils

const (
	SUCCESS           = 520
	ErrorMysql        = 998
	ErrorData         = 999
	ErrorUserExist    = 1000
	ErrorUserNoExist  = 1001
	ErrorUserPassWord = 1002
	ErrorToken        = 1003
)

var codeMsg = map[int]string{
	SUCCESS:           "操作成功",
	ErrorData:         "参数非法",
	ErrorUserExist:    "用户存在",
	ErrorMysql:        "数据库问题",
	ErrorUserNoExist:  "用户不存在",
	ErrorUserPassWord: "密码错误",
	ErrorToken:        "token异常",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
