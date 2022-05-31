CREATE DATABASE shoppingapi;
USE shoppingapi;
CREATE TABLE Customers (
	Customer_Id int NOT NULL AUTO_INCREMENT,
    Name varchar(100) NOT NULL,
    Email varchar(30) NOT NULL,
    Street varchar(100) NOT NULL,
    City varchar(30) NOT NULL,
    Zip varchar(10) NOT NULL,
    Phone varchar(15) NOT NULL,
    Password_Hash char(64) NOT NULL,
    Date_of_Birth date,
    PRIMARY KEY (Customer_Id)
);
ALTER TABLE Customers
MODIFY COLUMN Password_Hash char(64) NOT NULL;
CREATE TABLE Customer_Cart (
	Cart_Id int NOT NULL AUTO_INCREMENT,
    Customer_Id int NOT NULL,
    Total float NOT NULL,
    Completed bool NOT NULL,
    PRIMARY KEY (Cart_Id),
    FOREIGN KEY (Customer_Id) REFERENCES Customers(Customer_Id)
);
ALTER TABLE Customer_Cart
ADD COLUMN Completed bool NOT NULL;
CREATE TABLE Cart_Item (
	Cart_Item_Id int NOT NULL AUTO_INCREMENT,
    Cart_Id int NOT NULL,
    Product_Id int NOT NULL,
    Quantity int NOT NULL,
    PRIMARY KEY (Cart_Item_Id),
    FOREIGN KEY (Cart_Id) REFERENCES Customer_Cart(Cart_Id),
    FOREIGN KEY (Product_Id) REFERENCES Products(Product_Id)
);
CREATE TABLE Product_Category (
	Category_Id varchar(50) NOT NULL,
    Category_Title varchar(50) NOT NULL,
    PRIMARY KEY (Category_Id)
);
CREATE TABLE Products (
	Product_Id int NOT NULL AUTO_INCREMENT,
    Category_Id varchar(50) NOT NULL,
    Title varchar(100) NOT NULL,
    Quantity int NOT NULL,
    Price float NOT NULL,
    Weight_Gram int NOT NULL,
    Description varchar(1000),
    PRIMARY KEY (Product_Id),
    FOREIGN KEY (Category_Id) REFERENCES Product_Category(Category_Id)
);
CREATE TABLE Orders (
	Order_Id int NOT NULL AUTO_INCREMENT,
    Customer_Id int NOT NULL,
    Total float NOT NULL,
    Weight int NOT NULL,
    Street varchar(50) NOT NULL,
    City varchar(30) NOT NULL,
    Zip varchar(10) NOT NULL,
    Phone varchar(15) NOT NULL,
    Order_date date NOT NULL,
    Logistic varchar(10) NOT NULL,
    Tracking_Number varchar(20) NOT NULL,
    PRIMARY KEY (Order_Id),
    FOREIGN KEY (Customer_Id) REFERENCES Customers(Customer_Id)
);
CREATE TABLE Order_Item (
	Order_Item_Id int NOT NULL AUTO_INCREMENT,
    Order_Id int NOT NULL,
    Product_Id int NOT NULL,
    Quantity int NOT NULL,
    PRIMARY KEY (Order_Item_Id),
    FOREIGN KEY (Order_Id) REFERENCES Orders(Order_Id),
    FOREIGN KEY (Product_Id) REFERENCES Products(Product_Id)
);



