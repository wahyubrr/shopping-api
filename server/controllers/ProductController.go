package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	// MODELS
	"github.com/wahyubrr/synapsis-backend-test/models"
)

// Get all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Call models to get the data
	result, err := models.GetAllProducts()	// MODELS
	if err != nil {
		panic(err.Error())
	}
	// call a method Next from result object
	var products []Product
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
func GetProduct(w http.ResponseWriter, r *http.Request) {
	// Call models to get the data
	result, err:= models.GetProduct(mux.Vars(r)["id"])	// MODELS
	if err != nil {
		panic(err.Error())
	}
	// call a method Next from result object
	var products []Product
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
func GetCategory(w http.ResponseWriter, r *http.Request) {
	// Call models to get the data
	result, err := models.GetCategory(mux.Vars(r)["id"])	// MODELS
	if err != nil {
		panic(err.Error())
	}
	// call a method Next from result object
	var products []Product
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

// Add product to the database (OPTIONAL)
func AddProduct(w http.ResponseWriter, r *http.Request) {
	var products Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// create a database object which can be used to connect with database
	_, err = models.AddProduct(models.Product(products))	// MODELS

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
