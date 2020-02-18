create table public.orders (
  id serial not null,
  user_id integer references users(id),
  isPaid boolean default false not null,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  last_active_at timestamp without time zone,
  constraint orders_pk primary key (id)
);
create table public.products (
  id serial not null,
  user_id integer references users(id),
  price integer,
  "name" text not null,
  "description" text null,
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
create table public.users (
  id serial not null,
  username text not null,
  "name" text not null,
  bio text not null,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  last_active_at timestamp without time zone,
  constraint users_pk primary key (id)
);