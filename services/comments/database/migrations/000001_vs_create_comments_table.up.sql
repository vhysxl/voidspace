CREATE TABLE IF NOT EXISTS comments  (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,            
    post_id INT NOT NULL,     
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_user (user_id),
    INDEX idx_post (post_id)
);