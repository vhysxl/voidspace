ALTER TABLE post_likes DROP COLUMN deleted_at;

DROP INDEX IF EXISTS idx_post_likes_deleted_at;
