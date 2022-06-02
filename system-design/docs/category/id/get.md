# Get A Product From Category

Get a specific product with using the product_id.

**URL** : `/api/category/{id}`

**Method** : `GET`

**Auth required** : NO

## Success Response

**Code** : `200 OK`

**Content examples**

Showing a specific products from a category.

```json
{
  "product_id": 6,
  "category_id": "chain",
  "title": "KMC X11EL 11-speed Chain W/ Quick Link",
  "quantity": 23,
  "price": 54.4,
  "weight_gram": 256,
  "description": "Compatible with Shimano / Sram / Campagnolo. Asymmetrical Chamfering Avoids Interference. Added Chamfering for Smooth Operation. Impeccable Precision｜Phenomenal Shifting Performance."
}.....
```