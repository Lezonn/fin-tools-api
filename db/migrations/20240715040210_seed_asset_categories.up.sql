INSERT INTO asset_categories 
  (asset_category_name, created_at, updated_at) 
VALUES 
  ('Cash', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Bank Account', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Mutual Funds', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Bonds', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Stock', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Crypto', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW())),
  ('Other Assets', UNIX_TIMESTAMP(NOW()), UNIX_TIMESTAMP(NOW()));