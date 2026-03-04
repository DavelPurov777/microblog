-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE posts_lists (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT now(),
    likes INT,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts_lists(id) ON DELETE CASCADE,
    UNIQUE (user_id, post_id)
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS posts_lists;
DROP TABLE IF EXISTS likes;
