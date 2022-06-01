package controllers

import "github.com/golang-jwt/jwt/v4"

var jwtKey = []byte("Y3ljbGluZ2lzbXlob2JieXNpbmNlaXdhc2FraWQ=")

// Product struct
type Product struct {
	Product_Id  int     `json:"product_id"`
	Category_Id string  `json:"category_id"`
	Title       string  `json:"title"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	Weight_Gram int     `json:"weight_gram"`
	Description string  `json:"description"`
}

// Customer struct
type Customer struct {
	Customer_Id   int    `json:"customer_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Street        string `json:"street"`
	City          string `json:"city"`
	Zip           string `json:"zip"`
	Phone         string `json:"phone"`
	Password_Hash string `json:"password_hash"`
	Date_of_Birth string `json:"date_of_birth"`
}

// a struct to read the email and password
type Credentials struct {
	Customer_Id int    `json:"customer_id"`
	Email       string `json:"email"`
}

// a struct that will be encoded to a JWT
type Claims struct {
	Customer_Id int    `json:"customer_id"`
	Email       string `json:"email"`
	jwt.StandardClaims
}

// a struct for sending a token to a MODEL
type Customer_Token struct {
	Token string `json:"token"`
}

// a struct for displaying cart
type Cart struct {
	Cart_Product  Product `json:"cart_product"`
	Cart_Quantity int     `json:"quantity"`
}

// a struct for putting an item to a customer's cart
type InputCart struct {
	Product_Id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}
