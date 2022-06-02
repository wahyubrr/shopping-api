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

func Checkout(w http.ResponseWriter, r *http.Request) {
	// get data from the body
	var payload InputCheckout
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

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

	// call models.Getcart with decoded jwt token (customer_id)
	result, err := models.Checkout(strconv.Itoa(claims.Customer_Id), payload.Logistic) // MODELS
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
