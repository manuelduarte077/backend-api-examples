-- This file should undo anything in `up.sql`
DROP TRIGGER IF EXISTS set_updated_at;
DROP INDEX IF EXISTS idx_tasks_title;
DROP INDEX IF EXISTS idx_tasks_completed;
DROP TABLE IF EXISTS tasks;