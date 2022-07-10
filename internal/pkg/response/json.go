package response

import (
	"github.com/gin-gonic/gin"
	"haolinju.xianhaohan.com/common"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Time      int64       `json:"time"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

func JSON(ctx *gin.Context, data interface{}, code int) {
	writeStatusCode(ctx.Writer, code)

	// 状态码
	var statusCode = code
	if code > StatusCustomStart {
		statusCode = StatusOK
	}

	resp := Response{
		Status:    code,
		Message:   CodeText(code),
		Time:      time.Now().Unix(),
		RequestId: common.Trace(ctx),
		Data:      data,
	}
	ctx.JSON(statusCode, resp)
	return
}

func writeStatusCode(w http.ResponseWriter, code int) {
	header := w.Header()
	header.Set("status-code", strconv.FormatInt(int64(code), 10))
}
