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
			c.JSON(401, gin.H{"error": "Sesi tidak valid"})
			return
		}

		var input models.Cart
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Format tidak valid"})
			return
		}

		var existingCart models.Cart
		if err := db.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&existingCart).Error; err == nil {
			db.Model(&existingCart).Update("quantity", existingCart.Quantity+1)
			c.JSON(200, gin.H{"message": "Jumlah barang ditambahkan ke keranjang"})
			return
		}

		input.UserID = userID.(uint)
		input.Quantity = 1
		db.Create(&input)

		c.JSON(200, gin.H{"message": "Barang baru ditambahkan ke keranjang"})
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
