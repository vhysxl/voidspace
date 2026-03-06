ALTER TABLE user_follows
ADD CONSTRAINT chk_no_self_follow CHECK (user_id <> target_user_id);