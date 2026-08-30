package main

import (
	"gocommerce/configs"
	"gocommerce/handlers"
	"gocommerce/middlewares"
	"gocommerce/migrations"
	"gocommerce/seeders" // Pastikan import ini ditambahkan
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	seeders.Seed(db) // Data otomatis masuk ke MySQL saat server menyala

	router := gin.Default()

	// Rute Publik
	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	// Rute Terlindungi (Protected)
	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.POST("/transactions", handlers.CreateTransaction(db))
	// ... (Tambahkan rute CRUD lainnya yang kamu butuhkan di sini)

	router.Run(":5000")
}
