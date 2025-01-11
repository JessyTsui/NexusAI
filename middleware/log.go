package middleware

import (
	"fmt"
	"nexus-ai/constant"

	"github.com/gin-gonic/gin"
)

func SetupLog(server *gin.Engine) {
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var requestID string
		if param.Keys != nil {
			requestID = param.Keys[string(constant.RequestIDKey)].(string)
		}
		return fmt.Sprintf("[API]     %s | %s | %d | %12v | %15s | %5s | %s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			requestID,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
		)
	}))
}
