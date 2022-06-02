package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// GET ALL PRODUCTS IN A CUSTOMER'S CART
// get customer_id from Customer_Cart
// 	if Customer_Cart not found (0), then create one
// 	if present, continue
// then get all data from cart_item table with the customer_id
func Order(customer_id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}

	query := "SELECT Order_Id, Total, Weight_Gram, Order_Date, Logistic, Tracking_Number, Payment_Completed FROM Orders WHERE Customer_Id=" + customer_id

	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return result, err
}
