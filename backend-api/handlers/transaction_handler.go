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
			c.JSON(401, gin.H{"error": "Sesi tidak valid"})
			return
		}

		var carts []models.Cart
		if err := db.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
			c.JSON(500, gin.H{"error": "Gagal membaca keranjang"})
			return
		}
		if len(carts) == 0 {
			c.JSON(400, gin.H{"error": "Keranjang masih kosong"})
			return
		}

		tx := db.Begin()
		
		var grandTotal float64
		for _, cart := range carts {
			grandTotal += cart.Product.Price * float64(cart.Quantity)
		}

		transaction := models.Transaction{
			UserID: userID.(uint),
			Amount: grandTotal, 
		}
		if err := tx.Create(&transaction).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Gagal membuat transaksi"})
			return
		}

		for _, cart := range carts {
			txItem := models.TransactionItem{
				TransactionID: transaction.ID,
				ProductID:     cart.ProductID,
				Quantity:      uint(cart.Quantity),
				Price:         cart.Product.Price,
			}
			if err := tx.Create(&txItem).Error; err != nil {
				tx.Rollback()
				c.JSON(500, gin.H{"error": "Gagal mencatat rincian barang"})
				return
			}
		}

		if err := tx.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Gagal membersihkan keranjang"})
			return
		}

		tx.Commit()
		c.JSON(200, gin.H{"message": "Checkout sukses, transaksi tercatat secara permanen!"})
	}
}
