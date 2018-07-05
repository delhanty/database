-- +goose Up 
-- +goose NO TRANSACTION
CREATE DATABASE linkai WITH OWNER = linkai_admin TEMPLATE = template1 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8' TABLESPACE = pg_default CONNECTION LIMIT = -1;

-- +goose Down 
-- +goose NO TRANSACTION
-- SQL in this section is executed when the migration is rolled back.
DROP DATABASE linkai;
