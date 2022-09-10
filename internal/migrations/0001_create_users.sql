-- +migrate Up
CREATE TABLE users (
	user_id BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	user_name VARCHAR(100) not null,
    phone_number TEXT,
	is_active bool not null default true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX index_user_id ON users(user_id);
-- +migrate Down
DROP TABLE users;