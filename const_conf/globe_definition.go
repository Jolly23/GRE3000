package const_conf

import "encoding/json"

const (
	PlatformCookieName = "WX_OPENID_%s"
	OpenidCookieExpire = 60 * 60 * 24 * 365
)

const (
	Ok                  int = 0
	NotFound                = -1010
	JsonError               = -1011
	ApiNotFound             = -1012
	UserAccountNotFound     = -1013
	MenuKeyNotFound         = -1014
	PlatformNotFound        = -1015
	LogicError              = -1016

	ApiServerBusy      = -1
	ApiServerError     = -2
	CurrentServerError = -3
	ApiSignatureError  = -4
)

var ApiErrMsg = map[int]string{
	-5: "Request Error",
	-4: "API Signature Error",
	-3: "Current Server Error",
	-2: "API Server Error",
	-1: "API Server Busy",
	0:  "OK",
	1:  "Not Found",
	2:  "密码错误",
	10: "Unknown Error",

	-1010: "Not Found",
	-1011: "Json Error",
	-1012: "Api Not Found",
	-1013: "Current User's Account Not Found",
	-1014: "Menu Key Not Found",
	-1015: "Platform Not Found",
	-1016: "平台发生故障",
}

type ApiInfoBase struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Data    json.RawMessage
}

type BaseJsonReply struct {
	ErrCode int
	ErrMsg  string
}
