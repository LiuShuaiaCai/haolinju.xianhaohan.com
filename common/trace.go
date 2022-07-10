package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestId = "requestId"

func Trace(ctx *gin.Context) string {
	var requestId string

	if reqId, ok := ctx.Get(RequestId); ok {
		requestId = reqId.(string)
	} else {
		requestId = uuid.New().String()
		ctx.Set(RequestId, requestId)
	}

	return requestId
}
