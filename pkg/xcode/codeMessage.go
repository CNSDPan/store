package xcode

import "fmt"

var codeMessage = map[string]string{
	RESPONSE_SUCCESS:           "success",
	RESPONSE_FAIL:              "服务器内部错误",
	RESPONSE_NOT_FOUND:         "请求资源不存在",
	RESPONSE_UNAUTHORIZED:      "缺少身份认证",
	RESPONSE_NOT_CODE:          "无定义code码",
	RESPONSE_REQUEST_TIME_FAIL: "缺少请求日期",
	RESPONSE_TOKEN_FAIL:        "无效token",
	RESPONSE_APPID_FAIL:        "无效APPID",
	RESPONSE_SECRET_FAIL:       "无效secret",
	RESPONSE_SIGN_FAIL:         "无效sign",

	/***************************************/
	USER_LOGINED:            "用户已登录",
	USER_INFO_FAIL:          "用户信息不存在",
	USER_ID_FAIL:            "用户ID不存在|错误",
	USER_TOKEN_FAIL:         "用户Token不存在|错误",
	USER_TOKEN_CREATE:       "用户Token创建错误",
	USER_TOKEN_GET:          "用户Token获取错误",
	USER_TOKEN_DEL:          "用户Token删除错误",
	USER_CREAT_FAIL:         "用户创建失败",
	USER_CREAT_MOBILE_FAIL:  "用户手机号已创建",
	USER_CREAT_MOBILE_RULER: "手机号无效",
	USER_SET_INFOCACHE_FAIL: "用户信息存储缓存错误",
}

// GetCodeMessage
// @Desc：获取code码对应message内容
// @param：code
// @return：string
// @return：string
func GetCodeMessage(code string) (string, string) {
	var (
		message string
		ok      bool
		c       string
	)
	if message, ok = codeMessage[code]; !ok {
		message = fmt.Sprintf("原始Code:%s ,更改Code:%s ,message:%s", code, RESPONSE_NOT_CODE, codeMessage[RESPONSE_NOT_CODE])
		c = RESPONSE_NOT_CODE
	} else {
		c = code
	}
	return c, message
}