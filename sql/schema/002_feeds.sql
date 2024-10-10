-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE, 
    user_id UUID,
    CONSTRAINT fk_userid FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE 
);

-- +goose Down
DROP TABLE feeds;