CREATE TABLE IF NOT EXISTS event (
  id INT AUTO_INCREMENT PRIMARY KEY,
  type VARCHAR(255) NOT NULL,
  data TEXT NOT NULL,
  occurredAt DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
