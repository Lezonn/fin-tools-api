INSERT INTO expense_categories 
  (expense_category_name, created_at, updated_at) 
VALUES 
  ('Food & Beverage', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Transportation', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Fun Money', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Gifts & Donations', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Education', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Insurance', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Investment', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Pay Interest', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Houseware', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Personal Items', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Utility Bills', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Other Expenses', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW()));