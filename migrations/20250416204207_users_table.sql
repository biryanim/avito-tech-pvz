-- +goose Up
-- +goose StatementBegin

create table users
(
    id uuid primary key default gen_random_uuid(),
    email varchar(255) not null unique,
    role varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table users;

-- +goose StatementEnd
