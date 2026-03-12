ALTER TABLE post_like DROP COLUMN deleted_at;

DROP INDEX IF EXISTS idx_posts_deleted_at;