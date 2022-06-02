# Add a product to the store

Add a product to the store

**URL** : `/api/product`

**Method** : `POST`

**Auth required** : NO

**Data constraints**
```json
{
  "category_id": "[char 50 max]",
  "title": "[char 100 max]",
  "quantity": [int],
  "price": [int],
  "weight_gram": [int],
  "description": "[char 1000 max]"
}
```

**Data example**

```json
{
  "category_id": "chain",
  "title": "Shimano CN-HG93 Ultegra/XT 9-Speed Chain",
  "quantity": 13,
  "price": 29.42,
  "weight_gram": 304,
  "description": "The Shimano Deore XT CN-HG93 9-speed bicycle chain helps give your mountain bike smooth, precise shifting."
}
```

## Success Response

**Code** : `201 Created`