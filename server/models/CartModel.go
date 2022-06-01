package models

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// GET ALL PRODUCTS IN A CUSTOMER'S CART
// get customer_id, then use it for finding cart_id that is not yet completed
// 	if cart not foound (0), then create one
// 	if present, then get cart_id
// then get all data from cart_item table with the cart_id
func GetCart(customer_id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT Cart_Id FROM Customer_Cart WHERE Customer_Id=" +
		customer_id + " AND Completed=0"

	result, err = db.Query(query)
	var cart_id string
	// if no customer_cart is present for this customer
	if result.Next() == false {
		// fmt.Println("It is null :/")
		// create customer_cart
		query = "INSERT INTO Customer_Cart (Customer_Id, Total, Completed) VALUES (" +
			customer_id + ", 0, 0)"
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		// search the new cart_id of this customer
		query = "SELECT Cart_Id FROM Customer_Cart WHERE Customer_Id=" + customer_id
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		result.Next()
	}
	result.Scan(&cart_id)

	// search the founded cart_id in cart_item table
	query = "SELECT Products.Product_Id, Products.Category_Id, Products.Title, Products.Quantity, Products.Price, Products.Weight_Gram, Products.Description, Cart_Item.Quantity AS Cart_Quantity FROM Products INNER JOIN Cart_Item ON Products.Product_Id=Cart_Item.Product_Id WHERE Cart_Id=" + cart_id
	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return result, err
}

// PUT A PRODUCT IN CUSTOMER'S CART
// if customer_id == 7 && cart == 0 && completed == false (dont have cart) // check if customer_id 7 have a cart
// 	create Customer_cart
// else
// 	get cart_id

// // adding items
// add item to cart
// if customer_id == 7 && product_id == 10 (if cart_item is present)
// 	cart_item.quantity++
// else
// 	create Cart_item
func AddToCart(customer_id string, product_id string, quantity string) (result *sql.Rows, err error) {
	// fmt.Println(customer_id, product_id, quantity)
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	// GET CART_ID
	query := "SELECT Cart_Id FROM Customer_Cart WHERE Customer_Id=" +
		customer_id + " AND Completed=0"

	result, err = db.Query(query)
	var cart_id string
	// if customer doesn't have a customer_cart, create one
	if result.Next() == false {
		// fmt.Println("It is null :/")
		// create customer_cart
		query = "INSERT INTO Customer_Cart (Customer_Id, Total, Completed) VALUES (" +
			customer_id + ", 0, 0)"
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		// search the new cart_id of this customer
		query = "SELECT Cart_Id FROM Customer_Cart WHERE Customer_Id=" + customer_id
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		result.Next()
	}
	// cart_id is now available
	result.Scan(&cart_id)
	// fmt.Println(cart_id)

	// ADDING AN ITEM
	query = "SELECT Quantity FROM Cart_Item WHERE Cart_Id=" +
		cart_id + " AND Product_Id=" +
		product_id
	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var old_quantity int
	if result.Next() == false {
		// create cart_item
		query = "INSERT INTO Cart_Item (Cart_Id, Product_Id, Quantity) VALUES (" +
			cart_id + ", " + product_id + ", " + quantity + ")"
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		return
	}
	result.Scan(&old_quantity)

	// query = "INSERT INTO Cart_Item (Cart_Id, Product_Id, Quantity) VALUES (" +
	// cart_id + ", " + product_id + ", " + quantity + ")"
	// add the new quantity with the old quantity
	var new_quantity string
	temp, _ := strconv.Atoi(quantity)
	new_quantity = strconv.Itoa(temp + old_quantity)

	// fmt.Println(temp, old_quantity, new_quantity)

	query = "UPDATE Cart_Item SET Quantity=" + new_quantity + " WHERE Cart_Id=" + cart_id + " AND Product_Id=" + product_id

	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return result, err
}
