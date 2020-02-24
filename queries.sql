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
  p.*
from public.products p
join public.users u on p.user_id = u.id
  and u.username = $1
where
  p.slug = $2;
-- name: GetProducts :many
select
  *
from public.products
where
  id > $1
order by
  id asc
limit
  $2;
-- name: GetUserProducts :many
select
  p.*
from public.products p
inner join public.users u on p.user_id = u.id
  and u.username = $1
where
  p.id > $2
order by
  id asc
limit
  $3;
-- name: MyProducts :many
select
  *
from public.products
where
  user_id = $1
  and id > $2
order by
  id asc
limit
  $3;
-- name: GetUser :one
select
  *
from public.users
where
  id = $1;
-- name: CreateUser :one
insert into public.users (username, "name", bio, email, password)
values
  ($1, $2, $3, $4, $5) returning *;
-- name: UpdatePassword :exec
update public.users
set
  password = $1
where
  id = $2;
-- name: PublishProduct :one
update public.products
set
  isPablished = true
where
  id = $1 returning *;
-- name: GetUserByEmail :one
select
  id, email, password
from public.users
where
  email = $1;
-- name: ListUsers :many
select
  *
from public.users
where
  id = any($1 :: int[]);
-- name: ListProducts :many
select
  *
from public.products
where
  id  =any($1 :: int []);