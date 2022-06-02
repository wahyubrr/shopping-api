# View a customer's Cart

View a customer's cart along with its products

**URL** : `/api/cart`

**Method** : `GET`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content examples**

Showing all the products.

```json
{
  "total_price": 484.95,
  "total_weight": 1050,
  "cart_product": [
    {
      "cart_quantity": 5,
      "product": {
        "product_id": 1,
        "category_id": "rear-derailleur",
        "title": "Shimano ULTEGRA Short Cage RD-R8000-SS 11-Speed",
        "quantity": 5,
        "price": 96.99,
        "weight_gram": 210,
        "description": "Updated with a SHIMANO SHADOW RD design, the SHIMANO ULTEGRA 8000 rear derailleur offers a sleek appearance and optimized shifting performance. Covers: 11-25T to 11-30T."
      }
    }
  ]
}.....
```