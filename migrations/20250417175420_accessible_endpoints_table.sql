-- +goose Up
-- +goose StatementBegin

create table accesses
(
    id integer generated always as identity,
    endpoint_address text not null,
    method varchar(10) not null,
    role varchar(20) not null
);

insert into accesses(endpoint_address, method, role) values
    ('/pvz', 'POST', 'moderator'),
    ('/pvz', 'GET', 'moderator'),
    ('/pvz', 'GET', 'employee'),
    ('/pvz/:pvzId/close_last_reception', 'POST', 'employee'),
    ('/receptions', 'POST', 'employee'),
    ('/products', 'POST', 'employee'),
    ('/pvz/:pvzId/delete_last_product', 'POST', 'employee')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
