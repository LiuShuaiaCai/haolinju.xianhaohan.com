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

func JSON(ctx *gin.Context, data interface{}, err error) {
	code := TextCode(err)

	writeStatusCode(ctx.Writer, code.Code)

	// 状态码
	var statusCode = code.Code
	if code.Code > StatusCustomStart {
		statusCode = StatusOK
	}

	resp := Response{
		Status:    code.Code,
		Message:   code.Msg,
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
