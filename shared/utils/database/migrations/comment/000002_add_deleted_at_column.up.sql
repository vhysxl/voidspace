ALTER TABLE comments ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;

CREATE INDEX IF NOT EXISTS idx_comments_deleted_at ON comments(deleted_at);