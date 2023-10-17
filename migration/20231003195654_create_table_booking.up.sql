CREATE TABLE `bookings`(
    id varchar(255) NOT NULL PRIMARY KEY,
    booking_code varchar(12) UNIQUE NOT NULL,
    patient_id varchar(255) NOT NULL,
    schedule_id  INT NOT NULL,
    booking_date DATETIME  NOT NULL,
    state  varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_booking_patients FOREIGN KEY (patient_id) REFERENCES patients(id),
    CONSTRAINT fk_booking_schedules FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);
