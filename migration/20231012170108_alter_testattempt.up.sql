ALTER TABLE test_attempt
ADD diagnosis text AFTER score;
ALTER TABLE test_attempt
ADD ai_diagnosis FLOAT AFTER score;
ALTER TABLE test_attempt
ADD ai_probability FLOAT AFTER score;
ALTER TABLE test_attempt
ADD ai_accuracy FLOAT AFTER score;
ALTER TABLE test_attempt
ADD status VARCHAR(255) AFTER feedback;