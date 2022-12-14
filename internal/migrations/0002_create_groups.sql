-- +migrate Up
CREATE TABLE groups (
    description varchar(255),
    group_name VARCHAR(100) not null,
    group_id BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    is_active bool not null default true,
    created_at timestamp not null default now(),
	updated_at timestamp not null default now()
);
CREATE UNIQUE INDEX index_group_id ON groups(group_id);
-- +migrate Down
DROP TABLE groups;