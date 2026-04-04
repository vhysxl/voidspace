CREATE TYPE deletion_source AS ENUM ('user', 'lifecycle');
ALTER TABLE posts ADD COLUMN deletion_source deletion_source DEFAULT NULL;

-- future refactor