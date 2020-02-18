-- Migration: mytable
-- Created at: 2020-02-17 15:12:52
-- ====  UP  ====
BEGIN;
create table public.users (
  id serial not null,
  username text not null,
  name text not null,
  bio text not null,
  created_at timestamp without time zone,
  updated_at timestamp without time zone,
  last_active_at timestamp without time zone,
  constraint users_pk primary key (id)
);
create unique index user_email_unique ON public.users (username);
COMMIT;
-- ==== DOWN ====
BEGIN;
drop table public.users;
COMMIT;