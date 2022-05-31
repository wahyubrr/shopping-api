package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Get all products
func GetCart(customer_id string) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", database_endpoint)
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT Cart_Id FROM Customer_Cart WHERE Customer_Id=" +
	customer_id + " AND Completed=0"
	
	result, err = db.Query(query)
	var cart_id string
	// if no customer_cart is present for this customer
	if result.Next() == false {
		fmt.Println("It is null :/")
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

// Put a product to cart
func AddToCart(customer_id string, product_id string, quantity string) (result *sql.Rows, err error) {
	fmt.Println(customer_id, product_id, quantity)
	db, err := sql.Open("mysql", database_endpoint)
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
		fmt.Println("It is null :/")
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
	fmt.Println(cart_id)

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
		cart_id + ", " + product_id + ", 0)"
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		// search the new cart_id of this customer
		query = "SELECT Cart_Item_Id FROM Cart_Item WHERE Cart_Id=" + cart_id
		result, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		result.Next()
	}
	result.Scan(&old_quantity)
	fmt.Println(old_quantity)

	// // search the founded cart_id in cart_item table
	// query = "SELECT Products.Product_Id, Products.Category_Id, Products.Title, Products.Quantity, Products.Price, Products.Weight_Gram, Products.Description, Cart_Item.Quantity AS Cart_Quantity FROM Products INNER JOIN Cart_Item ON Products.Product_Id=Cart_Item.Product_Id WHERE Cart_Id=" + cart_id
	// result, err = db.Query(query)
	// if err != nil {
	// 	panic(err.Error())
	// }

	return result, err
}

// Get a single product
// func GetProduct(id string) (result *sql.Rows, err error) {
// 	db, err := sql.Open("mysql", database_endpoint)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	query := "SELECT * FROM Products WHERE Product_Id=" + id
// 	result, err = db.Query(query)

// 	return result, err
// }

// // Get all products in a category
// func GetCategory(id string) (result *sql.Rows, err error) {
// 	db, err := sql.Open("mysql", database_endpoint)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	query := "SELECT * FROM Products WHERE Category_Id='" + id + "'"
// 	result, err = db.Query(query)

// 	return result, err
// }

// // Add product to the database (OPTIONAL)
// func AddProduct(products Product) (result *sql.Rows, err error) {
// 	db, err := sql.Open("mysql", database_endpoint)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	query := "INSERT INTO Products (Category_Id, Title, Quantity, Price, Weight_Gram, Description) VALUES('" +
// 		products.Category_Id + "', '" +
// 		products.Title + "', " +
// 		strconv.Itoa(products.Quantity) + ", " +
// 		fmt.Sprint(products.Price) + ", " +
// 		strconv.Itoa(products.Weight_Gram) + ", '" +
// 		products.Description + "')"
	
// 	result, err = db.Query(query)
	
// 	return result, err
// }