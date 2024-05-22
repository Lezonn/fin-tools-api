create table users (
  id bigint primary key auto_increment not null,
  email varchar(255) not null unique,
  name varchar(255) not null,
  password varchar(255) null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  deleted_at timestamp null default current_timestamp
);