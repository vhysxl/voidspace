ALTER TABLE post_likes ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;

CREATE INDEX IF NOT EXISTS idx_post_likes_deleted_at ON post_likes(deleted_at);
