package controllers

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"crypto/sha256"

	// _ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/mux"

	// MODELS
	"github.com/golang-jwt/jwt/v4"
	"github.com/wahyubrr/synapsis-backend-test/models"
)

// var jwtKey = []byte("Y3ljbGluZ2lzbXlob2JieXNpbmNlaXdhc2FraWQ=")

// Registring customers
func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var customers Customer
	err := json.NewDecoder(r.Body).Decode(&customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	// hashing
	hashed := sha256.Sum256([]byte(customers.Email + customers.Password_Hash))
	customers.Password_Hash = hex.EncodeToString(hashed[:])

	// create a database object which can be used to connect with database
	result, err := models.RegisterCustomer(models.Customer(customers))	// MODELS

	// check if email is avaliable
	var isNotAvailable int
	result.Scan(&isNotAvailable)
	if isNotAvailable == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		resp := make(map[string]string)
		resp["message"] = "Status email already taken"
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp := make(map[string]string)
	resp["message"] = "Status Customer Created"
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	json.NewEncoder(w).Encode(resp)
}

// Login customers
func LoginCustomer(w http.ResponseWriter, r *http.Request) {
	var customers Customer
	// var result_customers []Customer
	err := json.NewDecoder(r.Body).Decode(&customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// hashing
	hashed := sha256.Sum256([]byte(customers.Email + customers.Password_Hash))
	customers.Password_Hash = hex.EncodeToString(hashed[:])

	// create a database object which can be used to connect with database
	result, err := models.LoginCustomer(models.Customer(customers))	// MODELS

	if err != nil {
		panic(err.Error())
	}

	// if email or password is wrong, do this
	if result.Next() == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		resp := make(map[string]string)
		resp["message"] = "Status Unauthorized"
		//jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		// w.Write(jsonResp)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// expiration time for token
	expirationTime := time.Now().Add(30 * time.Minute)

	// Fill customer credentials from Models
	var creds Credentials
	result.Scan(&creds.Customer_Id, &creds.Email)

	// create the JWT claims, which includes customer_id, email and expiry time
	claims := &Claims {
		Customer_Id: creds.Customer_Id,
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// declare the token used for signing and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	resp := make(map[string]string)
	resp["message"] = "Status Accepted"
	resp["token"] = tokenString
	// jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	// w.Write(jsonResp)
	json.NewEncoder(w).Encode(resp)
}