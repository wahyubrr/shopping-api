package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/wahyubrr/synapsis-backend-test/routes"
)

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
	// Load .env file
	err := godotenv.Load("database-credentials.env")
	if err != nil {
		log.Fatalf("dotenv Error. Err: %s", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
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
