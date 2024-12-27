package model

type Product struct {
	ID          	int     `json:"id"`
	Name        	string  `json:"name"`
	Description 	string  `json:"description"`
	Cost       		float64 `json:"cost"`
	Count			int     `json:"count"`
	ManufacturerId 	int 	`json:"manufacturerId"`
	SupplierId		int 	`json:"supplierId"`
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	AccessLevel string `json:"role"`
}