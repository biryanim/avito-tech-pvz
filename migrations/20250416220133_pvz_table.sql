-- +goose Up
-- +goose StatementBegin

create table pvz
(
    id uuid primary key,
    registration_date timestamp default now(),
    city varchar(50) not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table pvz;

-- +goose StatementEnd
