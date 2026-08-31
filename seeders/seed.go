package seeders

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt" // Tambahkan import bcrypt
)

func Seed(db *gorm.DB) {
	// Seeder untuk Product Categories
	category1 := models.ProductCategory{Name: "Category 1"}
	category2 := models.ProductCategory{Name: "Category 2"}

	db.Create(&category1)
	db.Create(&category2)

	// Proses hashing untuk seeder
	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte("password1"), bcrypt.DefaultCost)
	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("password2"), bcrypt.DefaultCost)

	// Seeder untuk Users menggunakan password yang sudah di-hash
	user1 := models.User{Username: "user1", Email: "user1@example.com", Password: string(hashedPassword1)}
	user2 := models.User{Username: "user2", Email: "user2@example.com", Password: string(hashedPassword2)}

	db.Create(&user1)
	db.Create(&user2)

	/// Seeder untuk Products
	product1 := models.Product{Name: "Product 1", CategoryID: 1}
	product2 := models.Product{Name: "Product 2", CategoryID: 2}

	db.Create(&product1)
	db.Create(&product2)

	// Seeder untuk Transactions
	transaction1 := models.Transaction{UserID: user1.ID, Amount: 100.0}
	transaction2 := models.Transaction{UserID: user2.ID, Amount: 200.0}

	db.Create(&transaction1)
	db.Create(&transaction2)

	// Seeder untuk Transaction Items
	item1 := models.TransactionItem{TransactionID: transaction1.ID, ProductID: product1.ID, Quantity: 2}
	item2 := models.TransactionItem{TransactionID: transaction2.ID, ProductID: product2.ID, Quantity: 3}

	db.Create(&item1)
	db.Create(&item2)
}
