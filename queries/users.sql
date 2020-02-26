
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
