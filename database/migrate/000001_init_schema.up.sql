CREATE TABLE IF NOT EXISTS urls (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,         -- 使用 BIGINT AUTO_INCREMENT
  original_url TEXT NOT NULL,
  short_code VARCHAR(255) NOT NULL UNIQUE,      -- 改为 VARCHAR(255)
  is_custom TINYINT(1) NOT NULL DEFAULT 0,      -- 使用 TINYINT(1) 表示布尔值
  expired_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_short_code ON urls (short_code);
CREATE INDEX idx_expired_at ON urls (expired_at);

   