package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 記錄請求開始的時間
		startTime := time.Now()

		// 處理請求
		c.Next()

		// 記錄請求所需的時間
		latency := time.Now().Sub(startTime)
		log.Printf("請求 %s - %s; 耗時: %v", c.Request.Method, c.Request.URL.Path, latency)
	}
}
