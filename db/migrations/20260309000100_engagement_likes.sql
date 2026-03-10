-- +goose Up
CREATE TABLE engagement_post_likes (
    post_id INT PRIMARY KEY,
    likes BIGINT NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE IF EXISTS engagement_post_likes;