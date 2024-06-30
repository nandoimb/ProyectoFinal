package main

import (
	"autmtres/models"
	"autmtres/repository"
	"autmtres/services"
	"strconv"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	productService services.ProductService
	orderService   services.OrderService
)

func initDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{}, &models.OrderProduct{}, &models.CartItem{})
	return db
}

func main() {
	r := gin.Default()
	db := initDatabase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	productRepo := repository.NewGormProductRepository(db)
	productService = services.NewProductService(productRepo)

	orderRepo := repository.NewGormOrderRepository(db)
	orderService = services.NewOrderService(orderRepo)

	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	r.GET("/orders", getOrders)
	r.GET("/orders/:id", getOrder)
	r.POST("/orders", createOrder)
	r.PUT("/orders/:id", updateOrder)
	r.DELETE("/orders/:id", deleteOrder)

	r.DELETE("/orders/:orderID/products/:productID", wrapRemoveProductFromOrder(db))

	r.Run(":8080")
}

// createProduct creates a new product
func createProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := productService.CreateProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, product)
}

// getProducts retrieves all products
func getProducts(c *gin.Context) {
	products, err := productService.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

// getProduct retrieves a product by ID
func getProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func updateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := productService.UpdateProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := productService.DeleteProduct(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}

// Order handlers
func createOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := orderService.CreateOrder(&order); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, order)
}

func getOrders(c *gin.Context) {
	orders, err := orderService.GetAllOrders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}

func getOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := orderService.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, order)
}

func updateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := orderService.UpdateOrder(&order); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}

func deleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := orderService.DeleteOrder(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}

// wrapRemoveProductFromOrder removes a product from an order
func wrapRemoveProductFromOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("orderID")
		productID := c.Param("productID")

		var orderProduct models.OrderProduct
		if err := db.Where("order_id = ? AND product_id = ?", orderID, productID).First(&orderProduct).Error; err != nil {
			c.JSON(404, gin.H{"error": "OrderProduct not found"})
			return
		}

		if err := db.Delete(&orderProduct).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(204, gin.H{})
	}
}
