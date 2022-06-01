package models

import (
	"fmt"
	"os"
	"strconv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Get all products
func GetAllProducts() (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT * FROM Products"
	result, err = db.Query(query)

	return result, err
}

// Get a single product
func GetProduct(id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT * FROM Products WHERE Product_Id=" + id
	result, err = db.Query(query)

	return result, err
}

// Get all products in a category
func GetCategory(id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT * FROM Products WHERE Category_Id='" + id + "'"
	result, err = db.Query(query)

	return result, err
}

// Add product to the database (OPTIONAL)
func AddProduct(products Product) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	query := "INSERT INTO Products (Category_Id, Title, Quantity, Price, Weight_Gram, Description) VALUES('" +
		products.Category_Id + "', '" +
		products.Title + "', " +
		strconv.Itoa(products.Quantity) + ", " +
		fmt.Sprint(products.Price) + ", " +
		strconv.Itoa(products.Weight_Gram) + ", '" +
		products.Description + "')"

	result, err = db.Query(query)

	return result, err
}
