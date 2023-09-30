ALTER TABLE users ADD specialization VARCHAR(255);
ALTER TABLE users ADD phone VARCHAR(255);
ALTER TABLE users ADD deleted_at TIMESTAMP DEFAULT NULL;
ALTER TABLE users RENAME COLUMN status TO state;
ALTER TABLE users RENAME COLUMN picture TO url_picture;
