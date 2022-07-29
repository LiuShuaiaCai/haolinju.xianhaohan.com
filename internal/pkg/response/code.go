package response

import (
	"sync"
)

const (
	StatusOK                  = 200 // RFC 7231, 6.3.1
	StatusBadRequest          = 400 // RFC 7231, 6.5.1
	StatusUnauthorized        = 401 // RFC 7235, 3.1
	StatusForbidden           = 403 // RFC 7231, 6.5.3
	StatusNotFound            = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed    = 405 // RFC 7231, 6.5.5
	StatusInternalServerError = 500 // RFC 7231, 6.6.1

	// 自定义错误码（从10000开始）
	StatusCustomStart   = 10000
	StatusCustomComment = 10001 // 自定义公共错误码
)

var codeText = map[int]string{
	StatusOK:                  "OK",
	StatusBadRequest:          "Bad Request",
	StatusUnauthorized:        "Unauthorized",
	StatusForbidden:           "Forbidden",
	StatusNotFound:            "Not Found",
	StatusMethodNotAllowed:    "Method Not Allowed",
	StatusInternalServerError: "Internal Server Error",

	// 自定义错误信息
}

var textCode = sync.Map{}

func init() {
	for code, msg := range codeText {
		textCode.Store(msg, code)
	}
}

func CodeText(code int) string {
	if message, ok := codeText[code]; ok {
		return message
	}

	return codeText[StatusInternalServerError]
}

type eCode struct {
	Code int
	Msg  string
}

func TextCode(err error) eCode {
	if err == nil {
		return eCode{
			Code: StatusOK,
			Msg:  codeText[StatusOK],
		}
	}
	msg := err.Error()
	var c int
	if code, ok := textCode.Load(msg); ok {
		c = code.(int)
	} else {
		c = StatusCustomComment
	}

	return eCode{
		Code: c,
		Msg:  msg,
	}
}
