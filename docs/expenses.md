# Expenses API
1. Add Expense

---

### Add Expense
> Used to add new expense.

+ Endpoint: **`/api/v1/expenses`**
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