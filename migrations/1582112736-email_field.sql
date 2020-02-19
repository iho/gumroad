-- Migration: email_field
-- Created at: 2020-02-19 13:45:36
-- ====  UP  ====

BEGIN;
ALTER TABLE public.users ADD COLUMN email text default '';
COMMIT;

-- ==== DOWN ====

BEGIN;
ALTER TABLE public.users DROP COLUMN email;
COMMIT;
