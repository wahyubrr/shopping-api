===API DOCS===
PRODUCTS
GET /product                       (Get all products)
GET /product/{id}                  (Get a product details)
GET /category/{id}   (Get all products from a category)
---bonus features---
POST /product               (Add new product - TOKEN and ADMIN)
PUT  /product/{id}          (Update product - TOKEN and ADMIN)
DELETE  /product/{id}       (Delete product - TOKEN and ADMIN)

POST /category              (Add new category - TOKEN and ADMIN)
PUT  /category/{id}   (Update category - TOKEN and ADMIN)
DELETE  /category/{id}       (Delete product - TOKEN and ADMIN)

CUSTOMERS
// GET     /customer/{id}  (Get customer by id)
GET     /customer       (Get all customer - with TOKEN)
POST    /customer       (Add new customer - with TOKEN)
PUT     /customer/{id}  (Update customer - with TOKEN)
DELETE  /customer/{id}  (Delete a customer - with TOKEN)

SHOPPING CART
GET     /cart       (Get a list of products in your shopping cart - with TOKEN)
PUT /cart/{id}      (Add to shopping cart - with TOKEN)
DELETE  /cart{id}   (Delete a particular product in your shopping cart - with TOKEN)

ORDERS
GET /orders/{id}    (Get a customer order - with TOKEN)
