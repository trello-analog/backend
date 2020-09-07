CREATE TABLE confirmation_codes (
    id bigserial not null primary key,
    user_id bigserial not null,
    code varchar not null,
    expired varchar not null
);
