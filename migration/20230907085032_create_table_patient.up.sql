CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255),
    nik VARCHAR(255),
    dob date,
    phone VARCHAR(20) NOT NULL,
    gender ENUM("male", "female"),
    marriage_status VARCHAR(100),
    nationality VARCHAR(255),
    partner_id VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    CONSTRAINT fk_patient_partner FOREIGN KEY (partner_id) REFERENCES patients(id)
);