package models

var database_endpoint string = "admin:Password8996290@tcp(synapsis-database-dev.c8oupwmizbjb.ap-southeast-1.rds.amazonaws.com:3306)/shoppingapi"

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

type Customer_Token struct {
	Token string `json:"token"`
}

// a struct for putting an item to a customer's cart
type InputCart struct {
	Product_Id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}