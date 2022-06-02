An Online Store API created with Go-lang with features including:
- Register and Login to an account
- View all product, and filter a product based on the category
- Customer can add and remove a product to/from their cart
- Customer can have more than one quantity to each product in their cart
- Customer can checkout their cart to order all selected items
- Customer then can pay for ther order (mockup)

# API Documentatioin

## Open Endpoints

Open endpoints require no Authentication.

* [Get All Products](system-design/docs/product/get.md) : GET /api/product

* [Get A Specific Product](system-design/docs/product/id/get.md) : GET /api/product/{id}

* [Get A Product From Category](system-design/docs/category/id/get.md) : GET /api/category/{category}

* [Add A Product to The Store](system-design/docs/product/post.md) : POST /api/product

* [Register an Account](system-design/docs/register/post.md) : POST /api/register

* [Login to an Account](system-design/docs/login/post.md) : POST /api/login

## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the request. A Token dan be acquired from the Login view above.

### Customer's Cart related

Each endpoint manipulates or displays information related to th customer's cart whose Token is provided with the request:

* [View Products from Customer's Cart](system-design/docs/cart/get.md) : GET /api/cart

* [Insert a Product to Customer's Cart](system-design/docs/cart/put.md) : PUT /api/cart

* [Delete a Products from Customer's Cart](system-design/docs/cart/delete.md) : DELETE /api/cart

### Checkout related

Endpoint for ordering all products from a customer's cart:

* [Check Out Customer's Cart](system-design/docs/checkout/post.md) : POST /api/checkout

### Payment related

Endpoint for paying a customer's cart (mockup):

* [Pay Customer's Order](system-design/docs/payment/post.md) : POST /api/payment

### Customer's Orders related

Endpoint for viewing a customers orders:

* [View Customer's Order](system-design/docs/order/get.md) : GET /api/order
