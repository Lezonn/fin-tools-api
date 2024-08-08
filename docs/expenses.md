# Expenses API
1. Create Expense
2. Delete Expense

---

### Create Expense
> Used to create new expense.

+ Endpoint: **`/api/expenses`**
+ HTTP Method: **`POST`**
+ Request Body:
```json5
{
  "expense_category_id": 1,
  "amount": 25000,
  "note": "Purchased fried chicken",
  "expense_date": 1689235200000
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

+ Response 404: `application/json`
```json5
{
  "code": 404,
  "status": "NOT_FOUND",
  "errors": {
    "errorMessage": ["Error message here"]
  }
}
```

+ Response 500: `application/json`
```json5
{
  "code": 500,
  "status": "INTERNAL_SERVER_ERROR",
  "errors": {
    "errorMessage": ["Error message here"]
  }
}
```
---

### Delete Expense
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

+ Response 500: `application/json`
```json5
{
  "code": 500,
  "status": "INTERNAL_SERVER_ERROR",
  "errors": {
    "errorMessage": ["Error message here"]
  }
}
```
---