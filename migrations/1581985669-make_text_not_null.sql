-- Migration: make_text_not_null
-- Created at: 2020-02-18 02:27:49
-- ====  UP  ====
BEGIN;
alter table public.products
alter column
  "description"
set
  default '',
alter column
  "description"
set
  not null;
alter table public.products
alter column
  summary
set
  default '',
alter column
  summary
set
  not null;
alter table public.products
alter column
  callToAction
set
  default '',
alter column
  callToAction
set
  not null;
alter table public.products
alter column
  coverImage
set
  default '',
alter column
  coverImage
set
  not null;
alter table public.products
alter column
  slug
set
  default '',
alter column
  slug
set
  not null;
alter table public.products
alter column
  receipt
set
  default '',
alter column
  receipt
set
  not null;
alter table public.products
alter column
  content
set
  default '',
alter column
  content
set
  not null;
-- ==== DOWN ====
  BEGIN;
alter table public.products
alter column
  "description" drop not null;
alter table public.products
alter column
  summary drop not null;
alter table public.products
alter column
  callToAction drop not null;
alter table public.products
alter column
  coverImage drop not null;
alter table public.products
alter column
  slug drop not null;
alter table public.products
alter column
  receipt drop not null;
alter table public.products
alter column
  content drop not null;
COMMIT;