# View Customer's Order

View a list of past order of this Customer

**URL** : `/api/order`

**Method** : `GET`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content examples**

Showing a list of past order of this customer.

```json
{
  {
        "order_id": 15,
        "total": 764,
        "total_weight": 247,
        "order_date": "2022-06-02 15:50:40",
        "logistic": "JNN",
        "tracking_number": "",
        "payment_completed": true
    },
    {
        "order_id": 16,
        "total": 96.99,
        "total_weight": 210,
        "order_date": "2022-06-02 15:52:23",
        "logistic": "JNN",
        "tracking_number": "7007353841",
        "payment_completed": true
    },....
  }
}
```