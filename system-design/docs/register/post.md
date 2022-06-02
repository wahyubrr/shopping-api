# Register An Account

Register a new account

**URL** : `/api/register`

**Method** : `POST`

**Auth required** : NO

**Data constraints**
```json
{
  "name": "[char 100 max]",
  "email": "[char 30 max]",
  "street": "[char 100 max]",
  "city": "[char 30 max]",
  "zip": "[char 10 max]",
  "phone": "[char 15 max]",
  "password": "[char 64 max]",
  "date_of_birth": [date]
}
```

**Data example**

```json
{
  "name": "Diana Nurlaily",
  "email": "diana@gmail.com",
  "street": "Jl. Bratang",
  "city": "Surabaya",
  "zip": "60115",
  "phone": "081330523131",
  "password": "takganti",
  "date_of_birth": "2001-08-21"
}
```

## Success Response

**Code** : `201 Created`