package main

import (
	"gocommerce/configs"
	"gocommerce/handlers"
	"gocommerce/middlewares"
	"gocommerce/migrations"
	_ "net/http/pprof"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	// seeders.Seed(db) // Data otomatis masuk ke MySQL saat server menyala

	router := gin.Default()

	// menambahkan Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Rute Publik
	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	// Rute Terlindungi (Protected)
	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.POST("/transactions", handlers.CreateTransaction(db))
	// menambahkan rute keranjang
	router.POST("/cart", middlewares.AuthMiddleware(), handlers.AddToCart(db))
	router.GET("/cart", middlewares.AuthMiddleware(), handlers.GetCart(db))
	router.POST("/checkout", middlewares.AuthMiddleware(), handlers.Checkout(db))

	router.Run(":5000")
}
