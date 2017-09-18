
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE comments(
	id SERIAL PRIMARY KEY,
	comments  TEXT NOT NULL,
	id_user SERIAL REFERENCES users(id),
	id_article SERIAL REFERENCES articles(id),
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE comments;
