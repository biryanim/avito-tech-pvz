-- +goose Up
-- +goose StatementBegin

create table pvz
(
    id uuid primary key,
    created_at timestamp default now(),
    city varchar(50) not null
);

create table receptions
(
    id uuid primary key,
    created_at timestamp not null default now(),
    pvz_id uuid not null references pvz(id) on delete cascade,
    status varchar(20) not null
);

create table products
(
    id uuid primary key,
    created_at timestamp not null default now(),
    type varchar(20) not null,
    reception_id uuid not null references receptions(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin


drop table products;
drop table receptions;
drop table pvz;

-- +goose StatementEnd
