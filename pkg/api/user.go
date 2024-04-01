package api

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "net/http"
    "myapp/model" // 引入用戶模型
    // 引入數據庫操作相關的包
)

// UpdateName 處理用戶名稱更新請求
func UpdateName(c *gin.Context) {
    session := sessions.Default(c)
    userID := session.Get("user_id")
    if userID == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "未登入或會話過期"})
        return
    }

    var newName struct {
        Name string `json:"name" binding:"required"`
    }
    if err := c.ShouldBindJSON(&newName); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 在這裡實現根據 userID 更新用戶名稱的邏輯
    // updateUserName(userID, newName.Name)

    c.JSON(http.StatusOK, gin.H{"message": "名稱更新成功"})
}
