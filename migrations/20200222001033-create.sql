-- +migrate Up
create table public.users (
  id serial not null,
  username text not null,
  "name" text not null,
  bio text not null,
  balance int default 0,
  password text default '',
  email text default '',
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  last_active_at timestamp without time zone,
  unique (username),
  constraint users_pk primary key (id)
);
-- +migrate Down
drop table public.users;