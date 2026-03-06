CREATE TABLE IF NOT EXISTS user_follows (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    target_user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_target_user
        FOREIGN KEY (target_user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT unique_follow UNIQUE (user_id, target_user_id)
);

CREATE INDEX idx_user ON user_follows(user_id);
CREATE INDEX idx_target_user ON user_follows(target_user_id);