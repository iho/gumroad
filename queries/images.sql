-- name: GetMyImages :many
select
  *
from public.images
where
  user_id = $1;
-- name: GetProductImages :many
select
  *
from public.images
where
  product_id = $1;