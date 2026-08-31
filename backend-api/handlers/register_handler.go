package handlers

import (
	"gocommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt" // Tambahkan import ini
)

// Pastikan struktur inputan disesuaikan dengan milikmu
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// --- PROSES ENKRIPSI BCRYPT MULAI DI SINI ---
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi kata sandi"})
			return
		}

		// Ganti password asli dengan password yang sudah di-hash
		input.Password = string(hashedPassword)
		// --------------------------------------------

		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat akun"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil"})
	}
}
