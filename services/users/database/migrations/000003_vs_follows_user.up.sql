CREATE TABLE user_follows (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,            
    target_user_id INT NOT NULL,     
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (target_user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE KEY unique_follow (user_id, target_user_id),
    INDEX idx_user (user_id),
    INDEX idx_target_user (target_user_id)
);