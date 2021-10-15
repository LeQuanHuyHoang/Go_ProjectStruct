package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LogFormatter(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("Log:\nClient ID:%s - TimeStamp:[%s] - Method:%s - Path:%s - Protocol:%s - StatusCode:%d - Latency:%s - UserAgent:%s - ErrerMessage:%s - x-request-id:%s\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
		param.Request.Header.Get("x-request-id"),
	)
}
