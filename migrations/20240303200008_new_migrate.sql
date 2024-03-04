-- +goose Up
create table users (
                      id serial primary key,
                      name varchar(50) not null,
                      email varchar(50) not null,
                      password varchar(50) not null,
                      role integer not null,
                      created_at timestamp not null default now(),
                      updated_at timestamp
);

-- +goose Down
drop table users;

