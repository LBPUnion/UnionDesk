-- Verify uniondesk:users-and-permissions on pg

BEGIN;

SELECT id, discord_id, avatar_url, "group", admin
FROM users
WHERE FALSE;

SELECT id, "group", "user", admin
FROM permission_bindings
WHERE FALSE;

ROLLBACK;
