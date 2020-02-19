-- Migration: password_field
-- Created at: 2020-02-19 13:40:28
-- ====  UP  ====

BEGIN;
ALTER TABLE public.users ADD COLUMN password text default '';

COMMIT;

-- ==== DOWN ====

BEGIN;
ALTER TABLE public.users DROP COLUMN password;
COMMIT;
