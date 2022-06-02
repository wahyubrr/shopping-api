package controllers

import (
	// "log"
	"encoding/json"
	"net/http"
	"strconv"

	// _ "github.com/go-sql-driver/mysql"
	// MODELS
	"github.com/golang-jwt/jwt/v4"
	"github.com/wahyubrr/synapsis-backend-test/models"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	// get the tokens from the header.Authorization
	temp_token := r.Header["Authorization"]
	// slice the token, because the raw file include "Bearer xxxxxxx", we remove the "Bearer "
	token := temp_token[0][7:] // this is token string

	// Initialize a new instance of "Claims"
	claims := &Claims{}

	// parse the token and fill the "claims" variable
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}

	// call models.Getcart with decoded jwt token (customer_id)
	result, err := models.GetCart(strconv.Itoa(claims.Customer_Id)) // MODELS
	if err != nil {
		panic(err.Error())
	}
	// call a method Next from result object
	var cart_one_product []Cart_One_Product
	var total_price float32
	var total_weight int
	for result.Next() {
		var product_id, quantity, weight_gram, cart_quantity int
		var category_id, title, description string
		var price float32
		err = result.Scan(&product_id, &category_id, &title, &quantity, &price, &weight_gram, &description, &cart_quantity, &total_price, &total_weight)
		if err != nil {
			panic(err.Error())
		}

		cart_one_product = append(cart_one_product, Cart_One_Product{
			Cart_Quantity: cart_quantity,
			Product: Product{
				Product_Id:  product_id,
				Category_Id: category_id,
				Title:       title,
				Quantity:    quantity,
				Price:       price,
				Weight_Gram: weight_gram,
				Description: description,
			},
		})
	}
	var completed_cart []Complete_Cart
	completed_cart = append(completed_cart, Complete_Cart{
		Cart_Total:   total_price,
		Cart_Weight:  total_weight,
		Cart_Product: cart_one_product,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed_cart)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	// get data from the body
	var input_cart InputCart
	err := json.NewDecoder(r.Body).Decode(&input_cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// fmt.Print(input_cart.Product_Id, input_cart.Quantity)

	// TOKEN
	// get the tokens from the header.Authorization
	temp_token := r.Header["Authorization"]
	// slice the token, because the raw file include "Bearer xxxxxxx", we remove the "Bearer "
	token := temp_token[0][7:] // this is token string
	// Initialize a new instance of "Claims"
	claims := &Claims{}
	// parse the token and fill the "claims" variable
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}

	// call models.Getcart with decoded jwt token (customer_id) and product_id
	result, err := models.AddToCart(strconv.Itoa(claims.Customer_Id), strconv.Itoa(input_cart.Product_Id), strconv.Itoa(input_cart.Quantity)) // MODELS
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	// get data from the body
	var input_cart InputCart
	err := json.NewDecoder(r.Body).Decode(&input_cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// fmt.Print(input_cart.Product_Id, input_cart.Quantity)

	// TOKEN
	// get the tokens from the header.Authorization
	temp_token := r.Header["Authorization"]
	// slice the token, because the raw file include "Bearer xxxxxxx", we remove the "Bearer "
	token := temp_token[0][7:] // this is token string
	// Initialize a new instance of "Claims"
	claims := &Claims{}
	// parse the token and fill the "claims" variable
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}

	// call models.Getcart with decoded jwt token (customer_id) and product_id
	result, err := models.DeleteFromCart(strconv.Itoa(claims.Customer_Id), strconv.Itoa(input_cart.Product_Id)) // MODELS
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
