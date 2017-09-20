
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE category(
	id SERIAL PRIMARY KEY,
	category character varying(25) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE comments;
