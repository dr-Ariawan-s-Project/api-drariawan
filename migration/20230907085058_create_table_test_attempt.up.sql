CREATE TABLE IF NOT EXISTS test_attempt (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    patient_id VARCHAR(255) NOT NULL UNIQUE,
    code_attempt VARCHAR(255),
    notes_attempt VARCHAR(255) ,
    score int,
    feedback TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    CONSTRAINT fk_patient_attempt FOREIGN KEY (patient_id) REFERENCES patients(id)
);