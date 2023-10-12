ALTER TABLE questions 
MODIFY type VARCHAR(255);
ALTER TABLE questions
MODIFY question text;
ALTER TABLE questions
ADD url_video text AFTER description;
ALTER TABLE questions
ADD section VARCHAR(255) AFTER url_video;
