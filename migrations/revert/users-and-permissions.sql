-- Revert uniondesk:users-and-permissions from pg

BEGIN;

DROP TABLE users;

DROP TABLE permission_bindings;

COMMIT;
