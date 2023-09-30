CREATE TABLE IF NOT EXISTS schedules (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    healthcare_address VARCHAR(255),
    day VARCHAR(255),
    time_start VARCHAR(5),
    time_end VARCHAR(5),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    CONSTRAINT fk_schedule_user FOREIGN KEY (user_id) REFERENCES users(id)
);