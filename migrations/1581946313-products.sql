-- Migration: products
-- Created at: 2020-02-17 15:31:53
-- ====  UP  ====
BEGIN;
create table public.products (
  id serial not null,
  user_id integer references users(id),
  price integer,
  name text not null,
  description text null,
  summary text null,
  callToAction text null,
  coverImage text null,
  slug text null,
  isPablished boolean default false not null,
  receipt text null,
  content text null,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  unique (user_id, slug),
  constraint products_pk primary key (id)
);
COMMIT;
-- ==== DOWN ====
BEGIN;
drop table public.products;
COMMIT;