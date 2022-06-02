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

// a struct for displaying completed cart
type Complete_Cart struct {
	Cart_Total   float32            `json:"total_price"`
	Cart_Weight  int                `json:"total_weight"`
	Cart_Product []Cart_One_Product `json:"cart_product"`
}

// a struct for displaying a single product in cart
type Cart_One_Product struct {
	Cart_Quantity int     `json:"cart_quantity"`
	Product       Product `json:"product"`
}

// a struct for putting an item to a customer's cart
type InputCart struct {
	Product_Id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type InputCheckout struct {
	Logistic string `json:"logistic"`
}

// Order struct
type Orders struct {
	Order_Id          int     `json:"order_id"`
	Total             float32 `json:"total"`
	Weight_Gram       int     `json:"total_weight"`
	Order_Date        string  `json:"order_date"`
	Logistic          string  `json:"logistic"`
	Tracking_Number   string  `json:"tracking_number"`
	Payment_Completed bool    `json:"payment_completed"`
}
