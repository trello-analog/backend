CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    login varchar not null unique,
    password varchar not null,
    two_auth boolean not null,
    avatar varchar
)