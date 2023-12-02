CREATE TABLE IF NOT EXISTS patient_token_passwords (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    patient_id VARCHAR(255),
    token TEXT,
    expired_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    CONSTRAINT fk_patients_token_passwords FOREIGN KEY (patient_id) REFERENCES patients(id)
);