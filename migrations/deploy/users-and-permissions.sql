-- Deploy uniondesk:users-and-permissions to pg

BEGIN;

CREATE TABLE users (
    id uuid primary key,
    discord_id text not null,
    avatar_url text,
    "group" uuid references groups(id),
    admin bool not null
);

CREATE TABLE permission_bindings (
    id uuid primary key,
    "group" uuid references groups(id),
    "user" uuid references users(id),

    admin bool not null
);

COMMIT;
