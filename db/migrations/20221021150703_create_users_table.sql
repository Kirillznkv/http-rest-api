-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists users (
   id bigserial not null primary key,
   email varchar not null unique,
   encrypted_password varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists users;
-- +goose StatementEnd
