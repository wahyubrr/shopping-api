package models

import (
	"database/sql"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// GET ALL PRODUCTS IN A CUSTOMER'S CART
// get customer_id from Customer_Cart
// 	if Customer_Cart not found (0), then create one
// 	if present, continue
// then get all data from cart_item table with the customer_id
func Payment(customer_id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}

	// random tracking number
	rand.Seed(time.Now().UnixNano())
	min, max := 1000000000, 9999999999
	tracking_number := strconv.Itoa(rand.Intn(max-min+1) + min)

	query := "UPDATE Orders SET Payment_Completed=1, Tracking_Number='" + tracking_number + "' WHERE Customer_Id=" + customer_id + " AND Payment_Completed=0;"

	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return result, err
}
