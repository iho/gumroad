-- Migration: add_balance
-- Created at: 2020-02-17 15:49:29
-- ====  UP  ====

BEGIN;
ALTER TABLE public.users ADD COLUMN balance INT default 0;
COMMIT;

-- ==== DOWN ====

BEGIN;
ALTER TABLE public.users DROP COLUMN balance;
COMMIT;
