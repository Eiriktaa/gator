-- +goose Up 
-- CREATE TABLE feed_follow (
--      id UUID PRIMARY KEY,
--     created_at TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP NOT NULL,
--     name TEXT NOT NULL,
--     url TEXT NOT NULL UNIQUE,
--     user_id UUID NOT NULL,
--     feed_id UUID NOT NULL,
--     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
--     FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
--     UNIQUE(user_id, feed_id)
--
-- );
-- +goose Down
DROP table if exists users cascade;
DROP TABLE if exists feeds cascade;
DROP TABLE if exists feed_follow cascade;
DROP TABLE if exists posts cascade;

