package handlers

import (
	"gocommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items").First(&transaction, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Transaction not found"})
			return
		}

		c.JSON(200, transaction)
	}
}

func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Transaction
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		for _, item := range input.Items {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				c.JSON(400, gin.H{"message": "Invalid product ID"})
				return
			}
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

func Checkout(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak valid"})
			return
		}

		// mengambil seluruh isi keranjang milik user
		var carts []models.Cart
		if err := db.Where("user_id = ?", userID).Find(&carts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca keranjang"})
			return
		}

		if len(carts) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Keranjang masih kosong"})
			return
		}

		// menghapus semua isi keranjang (simulasi checkout berhasil)
		if err := db.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses checkout"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pembayaran berhasil diproses! Terima kasih."})
	}
}
