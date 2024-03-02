-- +goose Up
create table user (
                      id serial primary key,
                      name text not null,
                      body text not null,
                      created_at timestamp not null default now(),
                      updated_at timestamp
);

-- +goose Down
drop table note;

