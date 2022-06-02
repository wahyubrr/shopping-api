# Get All Products

Get all products from the store.

**URL** : `/api/product`

**Method** : `GET`

**Auth required** : NO

## Success Response

**Code** : `200 OK`

**Content examples**

Showing all the products.

```json
{
  "product_id": 1,
  "category_id": "rear-derailleur",
  "title": "Shimano ULTEGRA Short Cage RD-R8000-SS 11-Speed",
  "quantity": 5,
  "price": 96.99,
  "weight_gram": 210,
  "description": "Updated with a SHIMANO SHADOW RD design, the SHIMANO ULTEGRA 8000 rear derailleur offers a sleek appearance and optimized shifting performance. Covers: 11-25T to 11-30T."
},
{
  "product_id": 2,
  "category_id": "rear-derailleur",
  "title": "SRAM RED eTap AXS Rear Derailleur 12-speed",
  "quantity": 4,
  "price": 764,
  "weight_gram": 247,
  "description": "The SRAM RED eTap AXS™ rear derailleur is AXS enabled for easy personalization and designed for both 1x and 2x systems. Advanced chain management keeps it silent and secure no matter the terrain. The derailleur also comes with faster pulleys, bearings, motors and signals, improving speed in every way."
}.....
```