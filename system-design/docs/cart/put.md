# Insert a product to a customer's cart

Insert a product to a customer's cart

**URL** : `/api/cart`

**Method** : `PUT`

**Auth required** : YES

**Data constraints**
```json
{
  "product_id": [int],
  "quantity": [int]
}
```

**Data example**

```json
{
  "product_id": 2,
  "quantity": 1
}
```

## Success Response

**Code** : `201 Created`