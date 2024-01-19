-- 20231125155534 - add_documents_table migration
PRAGMA journal_mode = wal;
PRAGMA busy_timeout = 5000;
PRAGMA synchronous = NORMAL;

CREATE TABLE documents (
   "id" TEXT PRIMARY KEY,
   "content" TEXT NOT NULL
);