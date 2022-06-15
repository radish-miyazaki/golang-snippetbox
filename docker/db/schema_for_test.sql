-- Create database for test
CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Create user for above database
CREATE USER test_web@127.0.0.1;
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'localhost';
ALTER USER test_web@127.0.0.1 IDENTIFIED BY 'pass';
