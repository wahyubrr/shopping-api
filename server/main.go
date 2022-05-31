package main

import (
	"fmt"

	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	//"math/rand"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var database_endpoint string = "admin:Password8996290@tcp(synapsis-database-dev.c8oupwmizbjb.ap-southeast-1.rds.amazonaws.com:3306)/shoppingapi"

// Product Model
type Product struct {
	Product_Id  int     `json:"product_id"`
	Category_Id string  `json:"category_id"`
	Title       string  `json:"title"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	Weight_Gram int     `json:"weight_gram"`
	Description string  `json:"description"`
}

// Get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	// create a database object which can be used to connect with database
	db, err := sql.Open("mysql", database_endpoint)
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT * FROM Products"
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	// call a method Next from result object
	for result.Next() {
		var product_id, quantity, weight_gram int
		var category_id, title, description string
		var price float32
		err = result.Scan(&product_id, &category_id, &title, &quantity, &price, &weight_gram, &description)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, Product{
			Product_Id:  product_id,
			Category_Id: category_id,
			Title:       title,
			Quantity:    quantity,
			Price:       price,
			Weight_Gram: weight_gram,
			Description: description,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get single product
func getProduct(w http.ResponseWriter, r *http.Request) {
	var products []Product
	// create a database object which can be used to connect with database
	db, err := sql.Open("mysql", database_endpoint)
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT * FROM Products WHERE Product_Id=" + mux.Vars(r)["id"]
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	// call a method Next from result object
	for result.Next() {
		var product_id, quantity, weight_gram int
		var category_id, title, description string
		var price float32
		err = result.Scan(&product_id, &category_id, &title, &quantity, &price, &weight_gram, &description)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, Product{
			Product_Id:  product_id,
			Category_Id: category_id,
			Title:       title,
			Quantity:    quantity,
			Price:       price,
			Weight_Gram: weight_gram,
			Description: description,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get all products in a category
func getCategory(w http.ResponseWriter, r *http.Request) {
	var products []Product
	// create a database object which can be used to connect with database
	db, err := sql.Open("mysql", database_endpoint)
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT * FROM Products WHERE Category_Id='" + mux.Vars(r)["id"] + "'"
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	// call a method Next from result object
	for result.Next() {
		var product_id, quantity, weight_gram int
		var category_id, title, description string
		var price float32
		err = result.Scan(&product_id, &category_id, &title, &quantity, &price, &weight_gram, &description)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, Product{
			Product_Id:  product_id,
			Category_Id: category_id,
			Title:       title,
			Quantity:    quantity,
			Price:       price,
			Weight_Gram: weight_gram,
			Description: description,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var products Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// create a database object which can be used to connect with database
	db, err := sql.Open("mysql", database_endpoint)
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

	_, err = db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// PRODUCTS Endpoints / Route Handlers
	r.HandleFunc("/api/product", getProducts).Methods("GET")
	r.HandleFunc("/api/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/category/{id}", getCategory).Methods("GET")

	// bonus features - CUD PRODUCTS
	r.HandleFunc("/api/product", addProduct).Methods("POST")
	// r.HandleFunc("/api/product/{id}", updateProduct).Methods("PUT")
	// r.HandleFunc("/api/product/{id}", deleteProduct).Methods("DELETE")

	// bonus features - CUD CATEGORY
	// r.HandleFunc("/api/category", addCategory).Methods("POST")
	// r.HandleFunc("/api/category/{id}", updateCategory).Methods("PUT")
	// r.HandleFunc("/api/product/{id}", deleteCategory).Methods("DELETE")

	db, err := sql.Open("mysql", "admin:Password8996290@tcp(synapsis-database-dev.c8oupwmizbjb.ap-southeast-1.rds.amazonaws.com:3306)/shoppingapi")
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

	log.Fatal(http.ListenAndServe(":8000", r))
}
