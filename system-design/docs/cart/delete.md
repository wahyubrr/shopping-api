# Delete a products from a customer's cart

Delete a products from a customer's cart

**URL** : `/api/cart`

**Method** : `DELETE`

**Auth required** : YES

**Data constraints**
```json
{
  "product_id": [int]
}
```

**Data example**

```json
{
  "product_id": 2
}
```

## Success Response

**Code** : `200 OK`