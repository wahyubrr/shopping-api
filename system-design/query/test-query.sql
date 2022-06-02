USE ShoppingDB;
SELECT * FROM Customers;
SELECT Email, Password_Hash FROM Customers WHERE Email="salsa@gmail.com" AND Password_Hash="takganti";
SELECT * FROM Customer_Cart;
UPDATE Customer_Cart SET Weight_Gram=210 WHERE Customer_Id=1;
SELECT * FROM Cart_Item;
SELECT * FROM Product_Category;
SELECT * FROM Products;
SELECT * FROM Products WHERE Category_Id="chain";
SELECT * FROM Orders;
SELECT * FROM Order_Item;

UPDATE Customer_Cart SET Weight_Gram=(SELECT CC.Weight_Gram - (P.Weight_Gram * CI.Quantity)
FROM (SELECT Weight_Gram, Customer_Id FROM Customer_Cart) AS CC, (SELECT Weight_Gram, Product_Id FROM Products) AS P,
	(SELECT Quantity, Product_Id FROM Cart_Item) AS CI
WHERE CC.Customer_Id=1 AND P.Product_Id=1 AND CI.Product_Id=1)
WHERE Customer_Cart.Customer_Id=1;

SELECT CC.Weight_Gram - (P.Weight_Gram * CI.Quantity)
FROM (SELECT Weight_Gram, Customer_Id FROM Customer_Cart) AS CC, (SELECT Weight_Gram, Product_Id FROM Products) AS P,
	(SELECT Quantity, Product_Id FROM Cart_Item) AS CI
WHERE CC.Customer_Id=1 AND P.Product_Id=1 AND CI.Product_Id=1;

UPDATE Customer_Cart SET Weight_Gram=(SELECT CC.Weight_Gram + (P.Weight_Gram * 2)
FROM (SELECT Weight_Gram, Customer_Id FROM Customer_Cart) AS CC, (SELECT Product_Id, Weight_Gram FROM Products) AS P
WHERE P.Product_Id=4 AND CC.Customer_Id=1)
WHERE Customer_Cart.Customer_Id=1;

SELECT CC.Weight_Gram + (P.Weight_Gram * 2)
FROM (SELECT Weight_Gram, Customer_Id FROM Customer_Cart) AS CC, (SELECT Product_Id, Weight_Gram FROM Products) AS P
WHERE P.Product_Id=4 AND CC.Customer_Id=1;

UPDATE Customer_Cart SET Weight_Gram=100 WHERE Customer_Id=1;

DELETE FROM Order_Item WHERE Order_Id=6;

INSERT INTO Order_Item (Order_Id, Product_Id, Quantity)
SELECT
	(SELECT Order_Id FROM Orders WHERE Order_Date=(
	SELECT MAX(Order_Date) FROM Orders
    )
),
Cart_Item.Product_Id, Cart_Item.Quantity FROM Cart_Item;

SELECT Order_Id FROM Orders WHERE Order_Date=(
	SELECT MAX(Order_Date) FROM Orders
);

SELECT MAX(Order_Date) AS DATES FROM Orders WHERE Customer_Id=1;

-- MODEL CHECKOUT
-- parameters: Customer_Id, Logistic, Tracking_Number(random from controller, mockup only)
-- needed column: Customer_Id, Customer_Cart.Total, Weight-tbd, Order_Date NOW(), Logistic, Tracking_Number
INSERT INTO Orders (Customer_Id, Total, Weight_Gram, Order_Date, Logistic, Tracking_Number)
	SELECT Customers.Customer_Id, Customer_Cart.Total, 1250 AS Weight_Gram,
	NOW() AS Order_Date, "JNE" AS Logistic, "jne12345" AS Tracking_Number
    FROM Customers, Customer_Cart WHERE Customers.Customer_Id=1 AND Customer_Cart.Customer_Id=1;

DELETE FROM Cart_Item WHERE Customer_Id=1;
DELETE FROM Customer_Cart WHERE Customer_Id=1;

DELETE FROM Cart_Item WHERE CUstomer_Id=1;
UPDATE Customer_Cart SET Total=0 WHERE Customer_Id=1;

(SELECT CC.Total - (P.Price * CI.Quantity) FROM (SELECT Total, Customer_Id FROM Customer_Cart) AS CC,
(SELECT Price, Product_Id FROM Products) AS P, (SELECT Quantity, Product_Id FROM Cart_Item) AS CI
WHERE CC.Customer_Id=1 AND P.Product_Id=6 AND CI.Product_Id=6);

UPDATE Customer_Cart SET Total=(SELECT CC.Total - (P.Price * CI.Quantity) FROM (SELECT Total, Customer_Id FROM Customer_Cart) AS CC,
(SELECT Price, Product_Id FROM Products) AS P, (SELECT Quantity, Product_Id FROM Cart_Item) AS CI
WHERE CC.Customer_Id=1 AND P.Product_Id=6 AND CI.Product_Id=6)
WHERE Customer_Cart.Customer_Id=1;


SELECT Products.Product_Id, Products.Category_Id, Products.Title, Products.Quantity, Products.Price, Products.Weight_Gram,
	Products.Description, Cart_Item.Quantity AS Cart_Quantity, Customer_Cart.Total AS Total
FROM ((Products
INNER JOIN Cart_Item ON Products.Product_Id=Cart_Item.Product_Id)
INNER JOIN Customer_Cart ON Cart_Item.Cart_Id=Customer_Cart.Cart_Id)
WHERE Cart_Item.Cart_Id=1;

-- Diketahui: customer_id dan product_id
-- objective: Customer_Cart.Total = Customer_Cart.Total - (Products.Price * Cart_Item.Quantity)
UPDATE Customer_Cart SET Total=(
	SELECT CC.Total - (P.Price * CI.Quantity) FROM (SELECT Total, Customer_Id FROM Customer_Cart) AS CC, (SELECT Price, Product_Id FROM Products) AS P, (SELECT Quantity, Product_Id FROM Cart_Item) AS CI
	WHERE CC.Customer_Id=1 AND P.Product_Id=2 AND CI.Product_Id=2
) WHERE Customer_Cart.Cart_Id=(
	SELECT CC.Cart_Id FROM (SELECT Cart_Id, Customer_Id FROM Customer_Cart) AS CC WHERE CC.Customer_Id=1
);

SELECT CC.Total - (P.Price * CI.Quantity) FROM (SELECT Total, Customer_Id FROM Customer_Cart) AS CC, (SELECT Price, Product_Id FROM Products) AS P, (SELECT Quantity, Product_Id FROM Cart_Item) AS CI
WHERE CC.Customer_Id=1 AND P.Product_Id=1 AND CI.Product_Id=1 ;

UPDATE Customer_Cart SET Total=0 WHERE Customer_Id=1;

SELECT Customer_Cart.Total + Products.Price FROM Customer_Cart, Products
WHERE Products.Product_Id=2;

-- Update the total cart
-- parameters: product_id, cart_id
UPDATE Customer_Cart SET Total=(
	SELECT CC.Total + (P.Price * 2) FROM (SELECT Total, Cart_Id FROM Customer_Cart) AS CC, (SELECT Product_Id, Price FROM Products) AS P
	WHERE P.Product_Id=8 AND CC.Cart_Id=2
)
WHERE Customer_Cart.Cart_Id=2;

-- CUSTOMER - REGISTER scratch code
-- parameters: email from Customers
INSERT INTO Customers (Name, Email, Street, City, Zip, Phone, Password_Hash, Date_of_Birth)
SELECT
	'Ash',
    'ash@gmail.com',
    'Jl. Kembar', 
    'Malang',
    '60333',
    '084123456',
    'takganti',
    '1996-02-01'
WHERE (
	SELECT COUNT(Email)FROM Customers WHERE Email='ash@gmail.com'
)=0;

SELECT COUNT(Email)FROM Customers WHERE Email='wahyubrlianto@gmail.com';

-- DELETECART scratch code
-- parameters: customer_id and product_id
DELETE FROM Cart_Item
WHERE Product_Id=6 AND Cart_Id=(
	SELECT Cart_Id FROM Customer_Cart
    WHERE Customer_Id=1
);

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