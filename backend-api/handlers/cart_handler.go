package handlers

import (
	"gocommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddToCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak valid"})
			return
		}

		var cart models.Cart
		if err := c.ShouldBindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data produk tidak valid"})
			return
		}

		cart.UserID = userID.(uint)
		cart.Quantity = 1

		if err := db.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memasukkan ke keranjang"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Berhasil ditambahkan ke keranjang"})
	}
}

func GetCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak valid"})
			return
		}

		var carts []models.Cart
		if err := db.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data keranjang"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": carts})
	}
}
