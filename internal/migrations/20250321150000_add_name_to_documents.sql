-- 20250321150000 - add_name_to_documents migration
PRAGMA journal_mode = wal;
PRAGMA busy_timeout = 5000;
PRAGMA synchronous = NORMAL;

ALTER TABLE documents ADD COLUMN name TEXT;
