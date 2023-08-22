CREATE TABLE IF NOT EXISTS questions (
    id INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    type VARCHAR(50) NOT NULL,
    question VARCHAR(255) NOT NULL,
    description TEXT,
    goto INT(11),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    CONSTRAINT fk_questions_goto_question FOREIGN KEY (goto) REFERENCES questions(id)
);

CREATE TABLE IF NOT EXISTS choices (
    id INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    question_id INT(11),
    `option` VARCHAR(255),
    score INT(11),
    slugs VARCHAR(255),
    goto INT(11),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    CONSTRAINT fk_questions_choices FOREIGN KEY (question_id) REFERENCES questions(id),
    CONSTRAINT fk_questions_goto_choice_question FOREIGN KEY (goto) REFERENCES questions(id)
);