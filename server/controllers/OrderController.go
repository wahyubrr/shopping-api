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

func Order(w http.ResponseWriter, r *http.Request) {
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
	order, err := models.Order(strconv.Itoa(claims.Customer_Id)) // MODELS
	if err != nil {
		panic(err.Error())
	}

	var orders []Orders
	for order.Next() {
		var total float32
		var order_id, weight_gram int
		var order_date, logistic, tracking_number string
		var payment_completed bool
		err = order.Scan(&order_id, &total, &weight_gram, &order_date, &logistic, &tracking_number, &payment_completed)
		if err != nil {
			panic(err.Error())
		}

		orders = append(orders, Orders{
			Order_Id:          order_id,
			Total:             total,
			Weight_Gram:       weight_gram,
			Order_Date:        order_date,
			Logistic:          logistic,
			Tracking_Number:   tracking_number,
			Payment_Completed: payment_completed,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
