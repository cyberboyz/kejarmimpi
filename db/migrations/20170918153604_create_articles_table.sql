
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE articles(
	id SERIAL PRIMARY KEY,
	title  TEXT NOT NULL,
	content  TEXT NOT NULL,
	id_user  SERIAL REFERENCES Users(id),
	id_category SERIAL REFERENCES category(id),
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE articles;
