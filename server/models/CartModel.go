package models

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// GET ALL PRODUCTS IN A CUSTOMER'S CART
// get customer_id from Customer_Cart
// 	if Customer_Cart not found (0), then create one
// 	if present, continue
// then get all data from cart_item table with the customer_id
func GetCart(customer_id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT Customer_Id FROM Customer_Cart WHERE Customer_Id=" +
		customer_id

	result, err = db.Query(query)

	// if no customer_cart is present for this customer, create one
	if result.Next() == false {
		// fmt.Println("It is null :/")
		// create customer_cart
		query = "INSERT INTO Customer_Cart (Customer_Id, Total) VALUES (" +
			customer_id + ", 0)"
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}

	}

	// search the customer_d in Cart_Item table
	query = "SELECT Products.Product_Id, Products.Category_Id, Products.Title, Products.Quantity, Products.Price, Products.Weight_Gram, Products.Description, Cart_Item.Quantity AS Cart_Quantity, Customer_Cart.Total AS Total, Customer_Cart.Weight_Gram FROM ((Products INNER JOIN Cart_Item ON Products.Product_Id=Cart_Item.Product_Id) INNER JOIN Customer_Cart ON Cart_Item.Customer_Id=Customer_Cart.Customer_Id) WHERE Cart_Item.Customer_Id=" + customer_id
	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return result, err
}

// PUT A PRODUCT IN CUSTOMER'S CART
// if customer_id == 7 && cart == 0 && completed == false (dont have cart) // check if customer_id 7 have a cart
// 	create Customer_cart

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
	// Search Customer_Cart is available
	query := "SELECT Customer_Id FROM Customer_Cart WHERE Customer_Id=" +
		customer_id

	result, err = db.Query(query)

	// if customer doesn't have a customer_cart, create one
	if result.Next() == false {
		// fmt.Println("It is null :/")
		// create customer_cart
		query = "INSERT INTO Customer_Cart (Customer_Id, Total) VALUES (" +
			customer_id + ", 0)"
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}

	}

	// Update the cart Total in Customer_Cart table
	// decrease Total in Cart if inputted quantity is negative
	check_quantity, err := strconv.Atoi(quantity)
	if err != nil {
		panic(err.Error())
	}
	if check_quantity < 0 {
		query = "UPDATE Customer_Cart SET Total=(SELECT CC.Total - (P.Price * " + quantity[1:] + ") FROM (SELECT Total, Customer_Id FROM Customer_Cart) AS CC, (SELECT Product_Id, Price FROM Products) AS P WHERE P.Product_Id=" + product_id + " AND CC.Customer_Id=" + customer_id + ")	WHERE Customer_Cart.Customer_Id=" + customer_id + ";"
	} else { // increase total in cart if quantity is positive
		query = "UPDATE Customer_Cart SET Total=(SELECT CC.Total + (P.Price * " + quantity + ") FROM (SELECT Total, Customer_Id FROM Customer_Cart) AS CC, (SELECT Product_Id, Price FROM Products) AS P WHERE P.Product_Id=" + product_id + " AND CC.Customer_Id=" + customer_id + ")	WHERE Customer_Cart.Customer_Id=" + customer_id + ";"
	}
	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// ADDING AN ITEM
	query = "SELECT Quantity FROM Cart_Item WHERE Customer_Id=" +
		customer_id + " AND Product_Id=" +
		product_id
	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var old_quantity int
	if result.Next() == false {
		// create cart_item
		query = "INSERT INTO Cart_Item (Customer_Id, Product_Id, Quantity) VALUES (" +
			customer_id + ", " + product_id + ", " + quantity + ")"
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		return
	}
	result.Scan(&old_quantity)

	// add the new quantity with the old quantity
	var new_quantity string
	inputted_quantity, _ := strconv.Atoi(quantity)
	new_quantity = strconv.Itoa(inputted_quantity + old_quantity)

	// fmt.Println(inputted_quantity, old_quantity, new_quantity)

	// Update the item quantity in Cart_item table
	query = "UPDATE Cart_Item SET Quantity=" + new_quantity + " WHERE Customer_Id=" + customer_id + " AND Product_Id=" + product_id
	result, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Update the cart's total weight

	return result, err
}

func DeleteFromCart(customer_id string, product_id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}
	query := "UPDATE Customer_Cart SET Total=(SELECT CC.Total - (P.Price * CI.Quantity) FROM (SELECT Total, Customer_Id FROM Customer_Cart) AS CC, (SELECT Price, Product_Id FROM Products) AS P, (SELECT Quantity, Product_Id FROM Cart_Item) AS CI WHERE CC.Customer_Id=" + customer_id + " AND P.Product_Id=" + product_id + " AND CI.Product_Id=" + product_id + ") WHERE Customer_Cart.Customer_Id=" + customer_id + ";"
	result, err = db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	query = "DELETE FROM Cart_Item WHERE Product_Id=" +
		product_id + " AND Customer_Id=" +
		customer_id

	result, err = db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	return result, err
}
