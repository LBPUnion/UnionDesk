-- Verify uniondesk:20240318-initial-migration on pg

BEGIN;

SELECT id, lang, content
FROM translations
WHERE FALSE;

SELECT id, role_id, name, description, ticket_category_id
FROM groups
WHERE FALSE;

SELECT id, parent, name, assignable
FROM categories
WHERE FALSE;

SELECT id, ticket_id, id_prefix, assigned_group, assignee, category, title
FROM tickets
WHERE FALSE;

SELECT id, ticket_id, content, sender
FROM comments
WHERE FALSE;

SELECT id, comment, hash, file_name, size
FROM attachments
WHERE FALSE;

ROLLBACK;
