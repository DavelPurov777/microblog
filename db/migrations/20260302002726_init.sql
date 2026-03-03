-- +goose Up
CREATE TABLE users (
    id serial not null unique,
    name varchar(255) not null ,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE posts_lists (
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    likes int
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS posts_lists;
