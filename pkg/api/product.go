package api

import (
	"/models"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	err := product.Create()
	if err != nil {
		// 處理錯誤
	}
	// 返回成功響應
}
