create table assets (
  id bigint auto_increment not null,
  user_id bigint not null,
  asset_category_id bigint not null default 0,
  amount bigint not null default 0,
  month int not null default 1,
  year int not null default 0,
  created_at bigint not null,
  updated_at bigint not null,
  primary key (id),
  foreign key (user_id) references users(id),
  foreign key (asset_category_id) references asset_categories(id)
)