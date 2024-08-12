# Expenses API
1. Create
2. Delete
3. Update
4. List
5. Get

---

### Create
> Used to create new expense.

+ Endpoint: **`/api/expenses`**
+ HTTP Method: **`POST`**
+ Request Body:
```json5
{
  "expense_category_id": 1,
  "amount": 25000,
  "note": "Purchased fried chicken",
  "expense_date": 1722330689
}
```

+ Response 200: `application/json`
```json5
{
  "code": 200,
  "status": "OK",
  "data": true
}
```

+ Response 500: `application/json`
```json5
{
  "code": 500,
  "status": "Internal Server Error",
  "message": "Error message here"
}
```
---

### Delete
> Used to delete expense.

+ Endpoint: **`/api/expenses/:id`**
+ HTTP Method: **`DELETE`**
+ Request Path :
  - id: `Integer`

+ Response 200: `application/json`
```json5
{
  "code": 200,
  "status": "OK",
  "data": true
}
```

+ Response 404: `application/json`
```json5
{
  "code": 404,
  "status": "Not Found",
  "message": "Error message here"
}
```

+ Response 500: `application/json`
```json5
{
  "code": 500,
  "status": "Internal Server Error",
  "message": "Error message here"
}
```
---

### Update
> Used to update expense.

+ Endpoint: **`/api/expenses/:id`**
+ HTTP Method: **`PATCH`**
+ Request Body:
```json5
{
  "expense_category_id": 2,
  "amount": 30000,
  "note": "Purchased fried chicken",
  "expense_date": 1722330689
}
```
+ Request Path :
  - id: `Integer`

+ Response 200: `application/json`
```json5
{
  "code": 200,
  "status": "OK",
  "data": true
}
```

+ Response 404: `application/json`
```json5
{
  "code": 404,
  "status": "Not Found",
  "message": "Error message here"
}
```

+ Response 500: `application/json`
```json5
{
  "code": 500,
  "status": "Internal Server Error",
  "message": "Error message here"
}
```
---

### List
> Used to get expense list.

+ Endpoint: **`/api/expenses`**
+ HTTP Method: **`GET`**
+ Response 200: `application/json`
```json5
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 1,
      "expense_category_id": 1,
      "amount": 30000,
      "note": "Purchased fried chicken",
      "expense_date": 1722330689
    },
    {
      "id": 2,
      "expense_category_id": 1,
      "amount": 10000,
      "note": "Purchased ice cream",
      "expense_date": 1722330689
    }
  ]
}
```
+ Response 500: `application/json`
```json5
{
  "code": 500,
  "status": "Internal Server Error",
  "message": "Error message here"
}
```
---

### Get
> Used to get expense detail by ID.

+ Endpoint: **`/api/expenses/:id`**
+ HTTP Method: **`GET`**
+ Request Path :
  - id: `Integer`
+ Response 200: `application/json`
```json5
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 1,
      "expense_category_id": 1,
      "amount": 30000,
      "note": "Purchased fried chicken",
      "expense_date": 1722330689
    }
  ]
}
```
+ Response 404: `application/json`
```json5
{
  "code": 404,
  "status": "Not Found",
  "message": "Error message here"
}
```
+ Response 500: `application/json`
```json5
{
  "code": 500,
  "status": "Internal Server Error",
  "message": "Error message here"
}
```
---