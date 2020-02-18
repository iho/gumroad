-- name: CreateProduct :one
insert into public.products (
    name,
    price,
    description,
    summary,
    callToAction,
    coverImage,
    slug,
    isPablished,
    receipt,
    content,
    user_id
  )
values
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
  ) returning *;
-- name: GetProduct :one
select
  name,
  price,
  description,
  summary,
  callToAction,
  coverImage,
  slug,
  isPablished,
  receipt,
  content
from public.products
where
  id = $1;
-- name: GetProducts :many
select
  p."name",
  p.price,
  p.description,
  p.summary,
  p.callToAction,
  p.coverImage,
  p.slug,
  p.isPablished,
  p.receipt,
  p.content
from public.products as p;
-- name: GetUser :one
select
  *
from public.users
where
  id = $1;
-- name: CreateUser :one
insert into public.users (username, "name", bio)
values
  ($1, $2, $3)
  returning *;