-- +goose Up
CREATE TABLE refresh_tokens (
    token text primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id uuid references users(id) on delete cascade not null,
    expires_at timestamp not null,
    revoked_at timestamp default null
);

-- +goose Down
DROP TABLE refresh_tokens;