-- +migrate Up
create table public.orders (
  id serial not null,
  user_id integer references users(id),
  isPaid boolean default false not null,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  last_active_at timestamp without time zone,
  constraint orders_pk primary key (id)
);
-- +migrate Down
drop table if exists public.orders;