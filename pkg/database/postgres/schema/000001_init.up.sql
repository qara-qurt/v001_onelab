CREATE TABLE users (
    id Serial PRIMARY KEY not null unique,
    fullName varchar(255) not null,
    login varchar(255) not null unique,
    password varchar(255) not null,
    isDeleted boolean default false
);

CREATE TABLE books (
    id Serial PRIMARY KEY not null unique,
    name varchar(255) not null,
    description varchar(255) not null,
    author varchar(255) not null
);

CREATE TABLE order_book_history (
    id Serial PRIMARY KEY not null unique,
    user_id INTEGER REFERENCES users(id) not null,
    book_id INTEGER REFERENCES books(id) not null,
    order_date TIMESTAMP NOT NULL DEFAULT NOW(),
    return_date TIMESTAMP
);

INSERT INTO users(fullName, login, password) VALUES('Dias Serikov','qara-qurt','qwerty');
INSERT INTO users(fullName, login, password) VALUES('test','test','qwerty');

INSERT INTO books(name, description, author) VALUES('Death note','Manga about Kira', 'test');
INSERT INTO books(name, description, author) VALUES('Harry Potter','Manga about Kira', 'J.K.ROWLING');
INSERT INTO books(name, description, author) VALUES('Зеленая миля','История об одной заложнике', 'Стивен Кинг');
INSERT INTO books(name, description, author) VALUES('Harry Potter 2','Мальчик который выжел', 'J.K.ROWLING');

INSERT INTO order_book_history(user_id, book_id, order_date, return_date) VALUES (1,1,CURRENT_TIMESTAMP,null);
INSERT INTO order_book_history(user_id, book_id, order_date, return_date) VALUES (1,2,CURRENT_TIMESTAMP,null);
INSERT INTO order_book_history(user_id, book_id, order_date, return_date) VALUES (2,3,CURRENT_TIMESTAMP,null);