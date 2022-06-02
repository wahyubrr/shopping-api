# Check Out Customer's Cart

Check out customer's cart

**URL** : `/api/checkout`

**Method** : `POST`

**Auth required** : YES

**Data constraints**
```json
{
  "logistic": "[char max 10]"
}
```

**Data example**

```json
{
  "logistic": "SUPER"
}
```

## Success Response

**Code** : `201 Created`

## Error Response

**Code** : "401 Unauthorized"

**Content example**

```json
{
  "Our stock is not enough for one of your ordered item"
}
```

```json
{
  "You have an unpaid order, Complete it first"
}
```