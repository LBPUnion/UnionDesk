-- Revert uniondesk:20240318-initial-migration from pg

BEGIN;

DROP TABLE attachments;
DROP TABLE comments;
DROP TABLE tickets;
DROP TABLE categories;
DROP TABLE groups;
DROP TABLE translations;

COMMIT;
