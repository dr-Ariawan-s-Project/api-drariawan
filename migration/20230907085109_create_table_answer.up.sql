CREATE TABLE IF NOT EXISTS answers (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    attempt_id VARCHAR(255),
    question_id INT,
    answer TEXT,
    score int,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    CONSTRAINT fk_attempt_answer FOREIGN KEY (attempt_id) REFERENCES test_attempt(id),
    CONSTRAINT fk_question_answer FOREIGN KEY (question_id) REFERENCES questions(id)
);