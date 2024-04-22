-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD CONSTRAINT email_unique UNIQUE (email);

create table resource_access (
    id serial primary key,
    role INTEGER not null,
    resource TEXT not null,
    is_access bool not null default false,
    CONSTRAINT unique_role_resource UNIQUE (role, resource)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP CONSTRAINT email_unique;
drop table resource_access;
-- +goose StatementEnd

