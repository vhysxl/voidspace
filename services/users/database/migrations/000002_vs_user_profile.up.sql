CREATE TABLE IF NOT EXISTS user_profile (
  user_id INT PRIMARY KEY,
  display_name VARCHAR(255),
  bio TEXT,
  avatar_url VARCHAR(500),
  banner_url VARCHAR(500),
  location VARCHAR(100),

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
