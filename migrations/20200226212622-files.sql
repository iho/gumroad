-- +migrate Up
create table public.files (
  id serial not null,
  user_id integer references users(id),
  product_id integer references products(id),
  path text,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  updated_at timestamp without time zone default (now() at time zone 'utc'),
  last_active_at timestamp without time zone default (now() at time zone 'utc'),
  constraint files_pk primary key (id)
);
-- +migrate Down
drop table if exists public.files;