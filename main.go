package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/golang-jwt/jwt"
	"time"
	"errors"
	"strconv"
)

const (
	ADMIN_ROLE = "Admin"
	USER_ROLE  = "User"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}


type Product struct {
	ID				string `json: "id"`
	Name 			string `json: "name"`
	Description 	string `json: "description"`
	Cost 			string `json: "cost"`
	Count 			string `json: "count"`
	ManufacturerId 	string `json: "manufacturerId"`
	SupplierId 		string `json: "supplierId"`
}

type User struct {
	Id       string
	Username string
	Password string
	Role     string
}

var products = []Product{
	{ID: "1", Name: "Продукт 1", Description: ":)", 	Cost: "560", 	Count: "10", ManufacturerId: "1", SupplierId: "1"},
	{ID: "2", Name: "Продукт 2", Description: "=)", 	Cost: "160", 	Count: "15", ManufacturerId: "1", SupplierId: "1"},
	{ID: "3", Name: "Продукт 3", Description: "О_О", 	Cost: "80", 	Count: "30", ManufacturerId: "2", SupplierId: "1"},
	{ID: "4", Name: "Продукт 4", Description: "-_-", 	Cost: "560", 	Count: "10", ManufacturerId: "2", SupplierId: "1"},
	{ID: "5", Name: "Продукт 5", Description: "xD", 	Cost: "2560", 	Count: "10", ManufacturerId: "3", SupplierId: "2"},
}

var users = []User{
	{
		Id:       "1",
		Username: "Admin",
		Password: "1234",
		Role:     "Admin",
	},
	{
		Id:       "2",
		Username: "user",
		Password: "1234",
	},
}



func main() {
	router := gin.Default()
	
	router.POST("/registrate", registrate)

	router.POST("/login", login)

	router.GET("/products", getProducts)

	router.GET("/products/:id", getProductsByID)

	router.GET("/users", authMiddleware(), getUsers)

	router.POST("/products", authMiddleware(), adminCheck(), createProduct)

	router.PUT("/products/:id", authMiddleware(), adminCheck(), updateProduct)

	router.DELETE("/products/:id", authMiddleware(), adminCheck(), deleteProduct)


	router.Run(":8080")
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductsByID(c *gin.Context) {
	id := c.Param("id")
	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}







func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	user, err := getUserByLogin(creds.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if creds.Username != user.Username || creds.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := generateToken(creds.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


func registrate(c *gin.Context) {
	var inputForm Credentials

	if err := c.BindJSON(&inputForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "input error"})
		return
	}

	if _, err := getUserByLogin(inputForm.Username); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage": "user already exists"})
		return
	}
	addUser(inputForm.Username, inputForm.Password)

	c.JSON(http.StatusCreated, gin.H{"message": "account created"})
}


func generateToken(username string, role string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)

		c.Next()
	}
}

func adminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		if role, exists := c.Get("role"); !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return

		} else if role != ADMIN_ROLE {
			c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to access this resource"})
			c.Abort()
		}

		c.Next()
	}
}



func getUserByLogin(login string) (User, error) {
	for _, user := range users {
		if user.Username == login {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
func getAllUsers() []User {
	return users
}

func addUser(login string, password string) {
	users = append(users, User{Id: strconv.Itoa(len(users) + 1), Username: login, Password: password, Role: USER_ROLE})
}

func getUsers(c *gin.Context) {
	users := getAllUsers()
	c.JSON(http.StatusOK, users)
}
