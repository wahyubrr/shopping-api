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
