create table users (
  id bigint primary key auto_increment not null,
  email varchar(255) not null unique,
  name varchar(255) not null,
  password varchar(255) null,
  google_id varchar(255) null unique,
  created_at bigint not null,
  updated_at bigint not null
);