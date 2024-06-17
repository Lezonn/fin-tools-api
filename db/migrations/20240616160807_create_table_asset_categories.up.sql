create table asset_categories (
  id bigint auto_increment not null,
  asset_category_name varchar(255) not null unique,
  created_at bigint not null,
  updated_at bigint not null,
  primary key (id)
)