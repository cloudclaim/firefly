BEGIN;
ALTER TABLE operations ADD COLUMN member VARCHAR(1024);
COMMIT;