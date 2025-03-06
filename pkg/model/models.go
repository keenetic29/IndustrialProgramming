package model

type Product struct {
	ID          	int    	`gorm:"primaryKey"`
	Name        	string  `json:"name"`
	Description 	string  `json:"description"`
	Cost       		float64 `json:"cost"`
	Count			int     `json:"count"`
	ManufacturerId 	int 	`json:"manufacturerId"`
	SupplierId		int 	`json:"supplierId"`
}

type User struct {
	ID 			int 	`gorm:"primaryKey"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
	AccessLevel string 	`gorm:"default:'User'"`
}