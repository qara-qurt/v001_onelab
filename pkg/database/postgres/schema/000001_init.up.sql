CREATE TABLE users (
    id Serial PRIMARY KEY not null unique,
    fullName varchar(255) not null,
    login varchar(255) not null unique,
    password varchar(255) not null
);

CREATE TABLE books (
    id Serial PRIMARY KEY not null unique,
    name varchar(255) not null,
    description varchar(255) not null,
    author_id INTEGER REFERENCES users(id)
);

INSERT INTO users(fullName, login, password) VALUES('Dias Serikov','qara-qurt','qwerty');