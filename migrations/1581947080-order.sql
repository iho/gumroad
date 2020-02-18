-- Migration: order
-- Created at: 2020-02-17 15:44:40
-- ====  UP  ====
BEGIN;
create table public.orders (
  id serial not null,
  user_id integer references users(id),
  isPaid boolean default false not null,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  last_active_at timestamp without time zone,
  constraint orders_pk primary key (id)
);
COMMIT;
-- ==== DOWN ====
BEGIN;
drop table public.orders;
COMMIT;