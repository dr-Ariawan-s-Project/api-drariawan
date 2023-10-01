ALTER TABLE users DROP COLUMN specialization;
ALTER TABLE users DROP COLUMN phone;
ALTER TABLE users DROP COLUMN deleted_at;
ALTER TABLE users RENAME COLUMN state TO status;
ALTER TABLE users RENAME COLUMN url_picture TO picture;