package routes

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	// CONTROLLERS
	"github.com/wahyubrr/synapsis-backend-test/controllers"
)

func StartController() {
	// Init Router
	r := mux.NewRouter()

	// PRODUCTS Endpoints / Route Handlers
	r.HandleFunc("/api/product", controllers.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/product/{id}", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/api/category/{id}", controllers.GetCategory).Methods("GET")

	// CART Endpoints / Route Handlers
	r.HandleFunc("/api/cart", controllers.GetCart).Methods("GET")
	r.HandleFunc("/api/cart", controllers.AddToCart).Methods("PUT")
	r.HandleFunc("/api/cart", controllers.DeleteFromCart).Methods("DELETE")

	// CUSTOMERS Endpoints / Route Handlers
	r.HandleFunc("/api/register", controllers.RegisterCustomer).Methods("POST")
	r.HandleFunc("/api/login", controllers.LoginCustomer).Methods("POST")

	// CHECKOUT Endpoints / Route Handlers
	r.HandleFunc("/api/checkout", controllers.Checkout).Methods("POST")

	// PAYMENT Endpoints / Route Handlers
	r.HandleFunc("/api/payment", controllers.Payment).Methods("POST")

	// Ordeer Endpoints / Route Handlers
	r.HandleFunc("/api/order", controllers.Order).Methods("GET")

	// bonus features - PRODUCTS - CUD PRODUCTS
	r.HandleFunc("/api/product", controllers.AddProduct).Methods("POST")
	// r.HandleFunc("/api/product/{id}", updateProduct).Methods("PUT")
	// r.HandleFunc("/api/product/{id}", deleteProduct).Methods("DELETE")

	// bonus features - CUD CATEGORY
	// r.HandleFunc("/api/category", addCategory).Methods("POST")
	// r.HandleFunc("/api/category/{id}", updateCategory).Methods("PUT")
	// r.HandleFunc("/api/product/{id}", deleteCategory).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
