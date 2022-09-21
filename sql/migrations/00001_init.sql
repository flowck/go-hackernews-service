-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(40) UNIQUE,
    email VARCHAR(100) UNIQUE,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now()
);

CREATE TABLE posts (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    title VARCHAR(240) NOT NULL,
    user_id UUID,
    votes INT,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
DROP TABLE users;
-- +goose StatementEnd
