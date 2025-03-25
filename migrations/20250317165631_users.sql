-- +goose Up
CREATE TABLE users (
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) unique not null,
    password varchar(255) not null,
    password_confirm varchar(255) not null, 
    role varchar(255) not null,
);
-- +goose Down
DROP TABLE users;
