package main

import (
	"gocommerce/configs"
	"gocommerce/handlers"
	"gocommerce/middlewares"
	"gocommerce/migrations"
	"gocommerce/models"
	"gocommerce/seeders"
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
	// seeders.Seed(db)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.POST("/transactions", handlers.CreateTransaction(db))

	router.POST("/cart", middlewares.AuthMiddleware(), handlers.AddToCart(db))
	router.GET("/cart", middlewares.AuthMiddleware(), handlers.GetCart(db))
	router.POST("/checkout", middlewares.AuthMiddleware(), handlers.Checkout(db))
	
	router.GET("/history", middlewares.AuthMiddleware(), func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		var history []models.Transaction

		db.Preload("Items").Preload("Items.Product").Where("user_id = ?", userID).Find(&history)

		c.JSON(200, gin.H{"data": history})
	})

	router.Run(":5000")
}
