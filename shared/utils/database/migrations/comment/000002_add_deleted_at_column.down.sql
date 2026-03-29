DROP INDEX IF EXISTS idx_comments_deleted_at;

ALTER TABLE comments DROP COLUMN IF EXISTS deleted_at;