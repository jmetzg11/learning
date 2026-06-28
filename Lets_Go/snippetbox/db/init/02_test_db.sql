-- Runs automatically the first time the MySQL container initializes (after
-- 01_schema.sql, by alphabetical order). Provisions an isolated database used
-- only by the Go integration tests (internal/models/*_test.go).
--
-- Only the database, user, and grant are created here. The test tables and
-- seed data are created/dropped on every test run by the testdata/setup.sql
-- and testdata/teardown.sql scripts, so the schema is intentionally absent.
--
-- The user is '%' (any host) rather than 'localhost' because `go test` runs on
-- the host and connects over TCP to the container, which MySQL sees as a
-- non-local connection.

CREATE DATABASE IF NOT EXISTS test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER IF NOT EXISTS 'test_web'@'%' IDENTIFIED BY 'pass';

GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'%';
