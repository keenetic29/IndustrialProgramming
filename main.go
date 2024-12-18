package main

import (
	"net/http"
	"time"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=1 dbname=api port=5432 sslmode=disable search_path=public"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Product_Order{})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
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
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
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
					c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
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
					c.JSON(http.StatusInternalServerError, gin.H{"message": "could not refresh token"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"token": newToken})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "token is still valid"})
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
	}
	router.Run(":8080")
}

func getBasket(c *gin.Context) {
	c.JSON(http.StatusOK, basket)
}

func addToBasket(c *gin.Context) {
	var newItem BasketItem

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
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
			c.JSON(http.StatusOK, gin.H{"message": "product removed from cart"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found in cart"})
}

func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.Model(&Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, updatedProduct)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
