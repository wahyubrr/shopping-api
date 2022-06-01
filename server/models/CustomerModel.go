package models

import (
	"database/sql"
	"os"
)

func RegisterCustomer(customers Customer) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}

	query := "SELECT COUNT(Email) FROM Customers WHERE Email='" +
		customers.Email + "'"
	result, err = db.Query(query)
	var isNotAvailable int
	result.Next()
	result.Scan(&isNotAvailable)
	if isNotAvailable == 1 {
		// email is already taken, so return isNotAvailable = 1 to controller
		return result, err
	}

	query = "INSERT INTO Customers (Name, Email, Street, City, Zip, Phone, Password_Hash, Date_of_Birth) VALUES ('" +
		customers.Name + "', '" +
		customers.Email + "', '" +
		customers.Street + "', '" +
		customers.City + "', '" +
		customers.Zip + "', '" +
		customers.Phone + "', '" +
		customers.Password_Hash + "', '" +
		customers.Date_of_Birth + "')"

	result, err = db.Query(query)

	return result, err
}

func LoginCustomer(customers Customer) (result *sql.Rows, err error) {
	db, err := sql.Open("mysql", os.Getenv("DB_WRITER_INSTANCE"))
	if err != nil {
		panic(err.Error())
	}

	query := "SELECT Customer_Id, Email FROM Customers WHERE " +
		"Email='" + customers.Email + "' AND Password_Hash='" +
		customers.Password_Hash + "'"

	result, err = db.Query(query)

	return result, err
}
