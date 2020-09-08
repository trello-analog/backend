CREATE TABLE forgot_password
(
    id        bigserial not null primary key,
    user_id   bigserial not null,
    code      varchar   not null,
    expired   varchar   not null,
    confirmed boolean   not null,
    last      boolean   not null
);
