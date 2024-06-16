create table expenses(
  id bigint auto_increment not null,
  user_id bigint not null,
  expense_category_id bigint not null default 0,
  amount bigint not null default 0,
  note text,
  created_at bigint not null,
  updated_at bigint not null,
  primary key (id),
  foreign key (user_id) references users(id),
  foreign key (expense_category_id) references expense_categories(id)
)