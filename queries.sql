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
  *
from public.products
where
  id = $1;
-- name: GetProducts :many
select
  *
from public.products
where
  user_id = $1 or $1 is null
  and id > $2
  order by id asc
  limit $3 ;
-- name: GetUser :one
select
  *
from public.users
where
  id = $1;
-- name: CreateUser :one
insert into public.users (username, "name", bio)
values
  ($1, $2, $3) returning *;
-- name: PublishProduct :one
update public.products
set
  isPablished = true
where
  id = $1 returning *;