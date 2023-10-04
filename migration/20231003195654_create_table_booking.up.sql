CREATE TABLE `bookings`(
    id INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    patient_id varchar(255) NOT NULL,
    schedule_id  INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_booking_patients FOREIGN KEY (patient_id) REFERENCES patients(id),
    CONSTRAINT fk_booking_schedules FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);
