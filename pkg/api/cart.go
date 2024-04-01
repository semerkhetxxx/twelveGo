package api

import (
	"/models"
	"net/http"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	cart := models.Cart{}
	err := cart.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create cart"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cart created successfully"))
}

func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	// Implement logic to add product to cart
}

func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	// Implement logic to remove product from cart
}
