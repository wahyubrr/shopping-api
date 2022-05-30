===API DOCS===
PRODUCTS
GET /products                       (Get all products)
GET /products/{id}                  (Get a product details)
GET /products/category/{category}   (Get all products from a category)
PUT /products/{id}                  (Add to shopping cart - with TOKEN)

CUSTOMERS
// GET     /customer/{id}  (Get customer by id)
GET     /customer       (Get all customer - with TOKEN)
POST    /customer       (Add new customer - with TOKEN)
PUT     /customer/{id}  (Update customer - with TOKEN)
DELETE  /customer/{id}  (Delete a customer - with TOKEN)

SHOPPING CART
GET     /cart       (Get a list of products in your shopping cart - with TOKEN)
DELETE  /cart{id}   (Delete a particular product in your shopping cart - with TOKEN)

ORDERS
GET /orders/{id}    (Get a customer order - with TOKEN)
