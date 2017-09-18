
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	profession TEXT NOT NULL,
	dream TEXT NOT NULL,
	hobby TEXT NOT NULL,
	website_url TEXT NOT NULL,
	phone TEXT DEFAULT NULL,
	avatar_url TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL,
	token TEXT NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
