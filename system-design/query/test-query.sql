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

INSERT INTO Product_Category (Category_Id, Category_Title) VALUES
("tire","Tire");

INSERT INTO Products (Category_Id, Title, Quantity, Price, Weight_Gram, Description) VALUES
("tire", "CONTINENTAL GRAND PRIX 5000 ROAD TIRE", 49, 82.95, 250, "Plenty of road tires on the market offer high performance, puncture resistance, or a supple ride, but very few can hit the sweet spot in multiple categories—that is, until the Grand Prix 5000. Continental’s newest edition of the famed Grand Prix line, the 5000 utilizes several different proprietary technologies in one tire, making some seriously capable rubber. Whether you’re racing, riding competitive group rides, or a weekend warrior, the Grand Prix 5000 is quite likely your ideal road tire."),
("tire", "PIRELLI P ZERO RACE 700C TIRE", 56, 79.9, 270, 'Pirelli worked closely with PRO cyclists to produce their high performing tubeless-ready road tire, the Pirelli P Zero Race 700c Tire. The tread pattern of the Pirelli P Zero Race 700c Tire is an evolution of the design of P ZERO™ Velo: several seasons of research and developement on the roads of the World Tour led to a new design that enhances the behaviour of the bicycle in every weather condition. The result is a race-day tyre with ultimate performance capabilities that can be used for everyday riding.');

INSERT INTO Products (Category_Id, Title, Quantity, Price, Weight_Gram, Description) VALUES
("pedal", "LOOK KEO CLASSIC 3 PLUS PEDALS", 25, 80, 350, "Clipless pedals have revolutionized the way riders ride their bikes. When you're clipped in, you're able to push and pull on the pedals for greater pedal performance and power transfer. The Look Keo Classic 3 Plus Pedals use a little extra stainless steel instead of composite for enhanced pedal stiffness and greater power transfer. The large and wide contact area provides excellent stability through the pedal stroke and at only 350g for the pair, these pedals will be great for any ride."),
("pedal", "LOOK KEO 2 MAX ROAD BIKE PEDALS", 23, 110, 280, "The Look team has completely remodeled the shape of the Keo 2 Max to give it a new, slimmed down line with a wider contact surface, increasing its style and efficiency. The width of the Look Keo 2 Max Road Pedal has been increased to 60mm to provide a larger platform for improved foot stability during pedaling. The 500 mm² of usable surface gives way for optimized power transfer, providing the same amount of contact for consistent performance. The spindle has also seen a few upgrades, starting with an oversized steel axle, which is composed of an axle with an inner ball bearing and a needle roller bearing. It’s located under the contact surface, allowing it to handle the load and distribute it optimally."),
("pedal", "LOOK KEO BLADE CARBON PEDALS", 14, 165, 230, "The type of pedals you use on your bike will have a significant impact on your pedaling performance and power. Carbon pedals are stiffer and lighter than aluminum and chromoly pedals, delivering more power to the rear wheel and providing a better ride quality for riders looking for higher performance. The Look Keo Blade Carbon Pedals boast both a carbon body and a carbon blade for exceptional performance. The pedals have a chromoly+ axle and stainless steel bearings for extreme reliability. With one of the best power-to-weight ratios, these pedals will be perfect for racers and riders looking to climb."),
("pedal", "SHIMANO 105 PD-R7000 SPD-SL BIKE PEDALS", 9, 150, 265, "Compliment the performance of your Shimano drivetrain by tying it together with the stellar performance of the 105 R7000 Pedal. The R7000 pedal delivers pro level performance for the cycling enthusiast. The extra-wide carbon reinforced platform provides more efficient power transfer, optimally converting your efforts to forward momentum. A wide bearing stance provides long lasting performance, and a buttery smooth action. At only 265g, the R7000 will not weigh down your rig."),
("pedal", "SHIMANO ULTEGRA PD-R8000 SPD-SL PEDALS", 11, 200, 248, "Compliment the performance of your Shimano drivetrain by tying it together with the stellar performance of the Ultegra R8000 Pedal. The R8000 pedal delivers pro level performance for the cycling enthusiast. The extra-wide carbon platform provides more efficient power transfer, optimally converting your efforts to forward momentum. A wide bearing stance provides long lasting performance, and a buttery smooth action. At only 248g, the R8000 will not weigh down your rig."),
("pedal", "SHIMANO DURA-ACE PD-R9100 SPD-SL PEDALS", 7, 280, 234, "These pedals feature an exceptionally wide and oversized bearing placement at both ends of the SPD-SL platform, which offers a symmetrical load dispersion and delivers a silky smooth feel with stable load distribution and superior pedaling performance. The low profile design provides maximum performance thanks to the close-to-axle pedal/cleat interface. Shimano is known to be very detailed oriented with their products, so it's no surprise to see when they design even the smallest of parts with superior performance in mind, like hollowed out cleats bolts for example, which are lightweight and durable. Ride to the top with R9100 Dura-Ace pedals!");

INSERT INTO Products (Category_Id, Title, Quantity, Price, Weight_Gram, Description) VALUES
("chain", "KMC X11SL TI NITRIDE 11-SPEED CHAIN", 23, 67.5, 239, "KMC respects the chain as perhaps the most key component on a bicycle. With this in mind they continuously innovate and strive to reach higher levels of performance from their chains. The KMC X11L Ti Nitride chain is a premium level 11 speed chain that delivers on performance, durability, reliability and happens to look good too. It has been tested and proven in the Beijing Olympics by USA cyclists in both road and mountain bike disciplines."),
("chain", "KMC X11EL TI NITRIDE 11-SPEED CHAIN", 36, 38.99, 256, "The KMC X11EL 11-speed Speed Chain is a quality replacement bicycle chain for Campagnolo, Shimano and SRAM drivetrain systems. Extra lightweight 11 speed chain featuring the Double-X Durability and X-Bridge technology. Double X Bridge Shape for super-fast and smooth shifting. The X11EL chain is compatible with all 11 speed road and MTB systems and is also non-directional in design. Fast and easy assembly of bike chain"),
("chain", "SRAM RED AXS CHAIN 12 SPEED", 20, 75, 250, "12-speed road drivetrains are finally here and with 12 gears on the rear means a new 12-speed chain. The SRAM Red AXS Chain is a narrower and lighter chain that has a unique link shape that increases the strength and provides more space between the outer plates of the chain and the cogs on either side. It features Flattop technology that enables a narrower chain with quieter operation and increased strength and durability. The Hard Chrome plated inner link plates and rollers provide reduced wear and prolong the life of the chain. This chain comes with 114 links and requires a new Flattop Powerlock."),
("chain", "SHIMANO CN-HG701 QL 11 SPEED CHAIN", 33, 38.99, 257, "The Shimano CN-HG701 QL 11 Speed Chain is a super Narrow - HYPERGLIDE - SIL-TEC -E-BIKE Chain that is E-bike compatible. The HG701-11 chain is directional and features SIL-TEC treatment on the roller link plate and pin link plate for increased durability. Th Shimano CN-HG701 QL 11 Speed Chain is compatible with 11-speed road, MTB, and E-BIKES thanks to its reinforced design. It’s also more efficient and silent with less maintenance needed thanks to the SIL-TEC surface treatment."),
("chain", "CONNEX 11SX 11-SPEED CHAIN", 12, 95.95, 270, "On the road or in the dirt, the Connex 11Sx 11-speed chain is made to excel in all terrain. The combination of high-grade nickel coating along with stainless steel creates a chain that is highly resistant to corrosion and is durable. The performance optimized plate design provides lightning fast shifting throughout the gear range. Complete with the Connex link for easy, tool free link connection.")
;

INSERT INTO Products (Category_Id, Title, Quantity, Price, Weight_Gram, Description) VALUES
("rear-derailleur", "SHIMANO DURA-ACE RD-R9250 DI2 REAR DERAILLEUR 12-speed", 10, 814.99, 215, "The new Dura-Ace RD-R9250 derailleur is the best way to experience Shimano’s latest and greatest 12-speed Di2 platform. It delivers Shimano’s fastest ever shifting with the new derailleur being a whopping 58% faster than the previous generation 11-speed Di2 derailleur. Its sleek and compact design has been refined while the cassette capacity can now support upt an 34t cassette. Straightforward charging and wireless connectivity make the install and setup process super easy. Being the Dura-Ace level derailleur this is the lightest possible derailleur available for your 12-speed Shimano road setup. So if you’re looking to shave off as many grams as possible then this is the way to go."),
("rear-derailleur", "SHIMANO ULTEGRA RD-R8000 11SP DERAILLEUR", 15, 96.99, 210, "Shimano has given the Ultegra 11spd an updated design that utilizes all the new features of the new dura Ace 9100 series at an affordable price. The Shimano Ultegra RD-R8000 Rear Derailleur uses Shimano's SHADOW RD low profile design, which reduces crash vulnerability while providing a sleek integrated appearance. There's two options available depending on the gear range you plan on using, for example the RD-SS covers 11-25T to 11-30T and the RD-GS covers 11-28T to 11-34T cassettes."),
("rear-derailleur", "SHIMANO 105 RD-R7000 REAR DERAILLEUR 11-speed", 11, 56.99, 232, "Shimano has redesigned the 105 line of components to keep it up to date with the ever-evolving industry standards. The Shimano 105 RD-R7000 Rear Derailleur utilizes Shimano’s low-profile Shadow parallelogram design, which reduces crash vulnerability and improves shifting performance. The new construction matches the force from the hand to reduce shifting effort and also allows for more flexible cable routing options. With the RD-GS option, you’ll be able to accommodate a 11-34t cassette, perfect for a gravel/CX rig."),
("rear-derailleur", "SRAM RED ETAP AXS REAR DERAILLEUR 12-speed", 6, 764, 247, "The heart of the connected cadre. The SRAM RED eTap AXS™ rear derailleur is AXS enabled for easy personalization and designed for both 1x and 2x systems. Advanced chain management keeps it silent and secure no matter the terrain. The derailleur also comes with faster pulleys, bearings, motors and signals, improving speed in every way."),
("rear-derailleur", "SRAM FORCE ETAP AXS REAR DERAILLEUR 12-speed", 17, 377, 325, "Fast, precise rear shifting is now more attainable than ever. Whether you ride 2x or 1x, tarmac or gravel, the SRAM Force eTap AXS rear derailleur is up for the task. It features Orbit damper technology for quiet, secure chain management over rough roads and capitalizes on larger X-SYNC pulleys for improved durability and efficiency."),
("rear-derailleur", "SRAM RIVAL ETAP AXS REAR DERAILLEUR 12-speed", 17, 377, 342, "On any surface you can imagine, in any drivetrain configuration, the Rival eTap AXS rear derailleur will handle it. Secure chain management and fast, precise shifting make for a smooth and quiet ride, every time.");

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