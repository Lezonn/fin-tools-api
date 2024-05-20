create table user_oauth_info (
  id bigint primary key auto_increment not null,
  oauth_provider varchar(20) not null,
  access_token varchar(40) not null,
  refresh_token varchar(40) not null,
  expiry_date timestamp not null
);