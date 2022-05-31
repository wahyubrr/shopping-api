package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/wahyubrr/synapsis-backend-test/routes"
)

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

func main() {
	db, err := sql.Open("mysql", database_endpoint)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Print("Successfully made the MySQL connection\n")

	routes.StartController()
}
