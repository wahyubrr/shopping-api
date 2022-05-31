SELECT * FROM Customers;
SELECT Email, Password_Hash FROM Customers WHERE Email="salsa@gmail.com" AND Password_Hash="takganti";
SELECT * FROM Customer_Cart;
SELECT * FROM Cart_Item;
SELECT * FROM Product_Category;
SELECT * FROM Products;
SELECT * FROM Products WHERE Category_Id="chain";
SELECT * FROM Orders;
SELECT * FROM Order_Item;

INSERT INTO Product_Category (
	Category_Id, Category_Title
)
VALUES (
	"chain", "Chain"
);
INSERT INTO Products (
	Category_Id,
    Title,
    Quantity,
    Price,
    Weight_Gram,
    Description
)
VALUES (
	"chain",
    "KMC X11EL 11-speed Chain W/ Quick Link",
    25,
    54.40,
    256,
    "Compatible with Shimano / Sram / Campagnolo. Asymmetrical Chamfering Avoids Interference. Added Chamfering for Smooth Operation. Impeccable Precision｜Phenomenal Shifting Performance."
);
DELETE FROM Products WHERE Product_Id=8;
UPDATE Products
SET Quantity=4, Weight_Gram=247, Title="SRAM RED eTap AXS Rear Derailleur 12-speed"
WHERE Product_Id=9;

INSERT INTO Customers (Name, Email, Street, City, Zip, Phone, Password_Hash, Date_of_Birth) VALUES (
	"Wahyu Berlianto",
    "wahyubrlianto@gmail.com",
    "Jl. Manyar Kertoarjo",
    "Surabaya",
    "60123",
    "081329443896",
    "takganti",
    "1999-02-01"
);
UPDATE Customers SET Password_Hash = "6a8fb09517fda8e6817ef0b74b3548b280cc67d827af1dd12d2b330dec33c9a4"
WHERE Customer_Id=1;