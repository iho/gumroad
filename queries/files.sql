-- name: GetMyFiles :many
select
  *
from public.files
where
  user_id = $1;
-- name: GetProductFile :one
select
  *
from public.files
where
  product_id = $1 limit 1;