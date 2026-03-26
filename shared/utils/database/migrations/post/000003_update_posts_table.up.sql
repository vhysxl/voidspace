ALTER TABLE posts ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;

CREATE INDEX IF NOT EXISTS idx_posts_deleted_at ON posts(deleted_at);