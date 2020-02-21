-- +migrate Up
create table public.products (
  id serial not null,
  user_id integer references users(id),
  price integer,
  "name" text not null,
  "description" text not null,
  summary text not null,
  callToAction text not null,
  coverImage text not null,
  slug text not null,
  isPablished boolean default false not null,
  receipt text not null,
  content text not null,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  unique (user_id, slug),
  constraint products_pk primary key (id)
);
-- +migrate Down
drop table public.users;