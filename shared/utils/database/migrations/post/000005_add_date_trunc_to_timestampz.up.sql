ALTER TABLE posts 
ALTER COLUMN created_at SET DEFAULT DATE_TRUNC('millisecond', NOW()),
ALTER COLUMN updated_at SET DEFAULT DATE_TRUNC('millisecond', NOW());