package main

import (
	"net/http"
	"twelveGo/config"
	"twelveGo/pkg/api"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.GetConfig()
	router := gin.Default()

	// 設置基於 Cookie 的會話
	store := cookie.NewStore([]byte("your_secret_key"))
	router.Use(sessions.Sessions("twelveGo_session", store))

	// 設置 HTML 模板
	router.LoadHTMLGlob("web/templates/*")

	// 路由設置
	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		// 假設您已經有方法從 JWT 中提取 claims 數據
		jwtClaims := GetJWTClaims(c) // 這需要您根據實際情況實現 GetJWTClaims 函數

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Port":      cfg.Port,
			"DBURL":     cfg.DBURL,
			"MongoURL":  cfg.MongoURL,
			"UserID":    userID,
			"JWTClaims": jwtClaims,
		})
	})

	router.GET("/heartbeat", api.Heartbeat)

	// 登入路由
	router.POST("/login", api.Login)

	// 受保護的路由
	router.GET("/protected", api.SomeProtectedRoute)

	router.GET("/memberheartbeat", api.MemberHeartbeat)

	// 自定義 404 處理
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 全局錯誤處理
	router.Use(func(c *gin.Context) {
		c.Next() // 處理請求

		// 檢查是否有任何錯誤
		if len(c.Errors) > 0 {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"errors": c.Errors,
			})
		}
	})

	router.Run(cfg.Port) // 啟動服務器
}
