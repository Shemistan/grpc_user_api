-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD CONSTRAINT email_unique UNIQUE (email);

create table url_access (
    id serial primary key,
    role INTEGER not null,
    url TEXT not null,
    is_access bool not null default false,
    CONSTRAINT unique_role_url UNIQUE (role, url)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP CONSTRAINT email_unique;
-- +goose StatementEnd

