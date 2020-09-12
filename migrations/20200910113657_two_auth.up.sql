CREATE TABLE two_auth
(
    id        bigserial not null primary key,
    user_id   bigserial not null references users(id),
    code      varchar   not null,
    expired   varchar   not null,
    confirmed boolean   not null,
    last      boolean   not null
);
