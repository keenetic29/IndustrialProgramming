package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

type Address struct {
	gorm.Model
	ID         uint
	CustomerID uint
	Country    string
	City       string
	Street     string
	House      int
}

type Customer struct {
	gorm.Model
	ID         uint
	Name       string
	Surname    string
	Patronymic string
	Telephone  string
	CardNumber string
	Mail       string
	Password   string
	Addresses  []Address
}

type Product struct {
	Id                string
	ManufacturerId    int
	ProductCategoryId int
	Name              string
	Price             float64
	Quantity          int
	ArticleNumber     int
	Description       string
	ImageUrl          string
}

type Product_Order struct {
	Id        string
	ProductId int
	OrderId   int
}

type Category struct {
	gorm.Model
	ID   uint
	Name string
}

type Manufacturer struct {
	gorm.Model
	ID      uint
	Name    string
	Country string
}

type Order struct {
	gorm.Model
	ID         uint
	Date       time.Time
	Time       time.Time
	CustomerID uint
	Deliveries []Delivery
	Products   []Product
}

type Review struct {
	gorm.Model
	ID         uint
	CustomerID uint
	ProductID  uint
	Date       time.Time
	Text       string
	Rating     int
}

type Favorite struct {
	gorm.Model
	ID         uint
	ProductID  uint
	CustomerID uint
}

type Provider struct {
	gorm.Model
	ManufacturerID uint
	ProductID      uint
}

type Basket struct {
	gorm.Model
	ID         uint
	CustomerID uint
	Products   []Product
}

type Delivery struct {
	gorm.Model
	ID           uint
	OrderID      uint
	AddressID    uint
	Status       string
	SendDate     time.Time
	ExpectedDate time.Time
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=1 dbname=api port=5432 sslmode=disable search_path=public"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	db.AutoMigrate(&Address{}, &Customer{}, &Product{}, &Order{}, &Review{}, &Favorite{}, &Category{}, &Manufacturer{}, &Provider{}, &Basket{}, &Delivery{})
	if err != nil {
		log.Fatalf("Не удалось выполнить миграцию базы данных: %v", err)
	}
}

var router = gin.Default()

var basket = []BasketItem{}

type BasketItem struct {
	ProductId string `json:"productId"`
	Quantity  int
}

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var users = []Credentials{
	{Username: "user", Password: "password"},
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
	{Username: "user3", Password: "password3"},
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	var validUser *Credentials
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			validUser = &user
			break
		}
	}

	if validUser == nil {
		handleError(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Could not create token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					handleError(c, http.StatusUnauthorized, "token expired")
					c.Abort()
					return
				}
			}
			handleError(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		if !token.Valid {
			handleError(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}

func refreshToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				newToken, err := generateToken(claims.Username)
				if err != nil {
					handleError(c, http.StatusInternalServerError, "could not refresh token")
					return
				}
				c.JSON(http.StatusOK, gin.H{"token": newToken})
				return
			}
		}
		handleError(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	if !token.Valid {
		handleError(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	handleError(c, http.StatusBadRequest, "token is still valid")
}

func main() {
	initDB()
	router.POST("/login", login)
	router.POST("/refresh", refreshToken)
	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		// Получение всех товаров
		router.GET("/products", getProducts)
		// Получение товара по ID
		router.GET("/products/:id", getProductByID)
		// Создание нового товара
		router.POST("/products", createProduct)
		// Обновление существующего товара
		router.PUT("/products/:id", updateProduct)
		// Удаление товара
		router.DELETE("/products/:id", deleteProduct)

		// Получение всех товаров в корзине
		protected.GET("/basket", getBasket)

		// Добавление товара в корзину
		protected.POST("/basket", addToBasket)

		// Удаление товара из корзины
		protected.DELETE("/basket/:productId", deleteFromBasket)

		protected.GET("/productswithtimeout", getProductsWithTimeout)
	}
	router.Run(":8080")
}

func getBasket(c *gin.Context) {
	c.JSON(http.StatusOK, basket)
}

func addToBasket(c *gin.Context) {
	var newItem BasketItem

	if err := c.BindJSON(&newItem); err != nil {
		handleError(c, http.StatusBadRequest, "invalid request")
		return
	}

	for i, item := range basket {
		if item.ProductId == newItem.ProductId {
			basket[i].Quantity += newItem.Quantity
			c.JSON(http.StatusOK, basket[i])
			return
		}
	}

	basket = append(basket, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func deleteFromBasket(c *gin.Context) {
	productId := c.Param("productId")

	for i, item := range basket {
		if item.ProductId == productId {
			basket = append(basket[:i], basket[i+1:]...)
			handleError(c, http.StatusOK, "product removed from basket")
			return
		}
	}
	handleError(c, http.StatusNotFound, "product not found in basket")

}

func getProducts(c *gin.Context) {
	var products []Product
	var total int64
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")
	name := c.Query("name")
	category := c.Query("category")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	offset := (pageInt - 1) * limitInt
	query := db.Limit(limitInt).Offset(offset)

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if category != "" {
		query = query.Where("category ILIKE ?", "%"+category+"%")
	}

	query.Find(&products).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": total,
		"page":  pageInt,
		"limit": limitInt,
	})
	sort := c.Query("sort")
	if sort != "" {
		query = query.Order(sort)
	}

}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		handleError(c, http.StatusNotFound, "Product not found")
		return
	}
	db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		handleError(c, http.StatusNotFound, "Product not found")
		return
	}
	if err := db.Model(&Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		handleError(c, http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, updatedProduct)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Product{}, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func getProductsWithTimeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	var products []Product
	if err := db.WithContext(ctx).Find(&products).Error; err != nil {
		handleError(c, http.StatusRequestTimeout, "Request timed out")
		return
	}

	c.JSON(http.StatusOK, products)
}
