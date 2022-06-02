# Login to An Account

Log in to a customer's account

**URL** : `/api/login`

**Method** : `POST`

**Auth required** : NO

**Data constraints**
```json
{
  "email": "[char 30 max]",
  "password": "[char 64 max]"
}
```

**Data example**

```json
{
  "email": "wahyubrlianto@gmail.com",
  "password": "takganti"
}
```

## Success Response

**Code** : `201 Created`

**Content example**

```json
{
  "message": "Status Accepted",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21lcl9pZCI6MSwiZW1haWwiOiJ3YWh5dWJybGlhbnRvQGdtYWlsLmNvbSIsImV4cCI6MTY1NDIwMTA3Mn0._pEYUqseTyQWR_wQhrLlSMo0B6x8DkQs6u1Q0GqnXKY"
}
```