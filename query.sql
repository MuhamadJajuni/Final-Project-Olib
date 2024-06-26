CREATE EXTENSION "uuid-ossp";

CREATE TABLE borrower (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name varchar(50) NULL,
    email varchar(100) NULL UNIQUE,
    password varchar(100) NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    address varchar(100) NULL
);

CREATE TABLE book (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title varchar(50) NULL,
    author varchar(100) NULL,
    release_year varchar(4) NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE admin (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name varchar(50) NULL,
    email varchar(100) NULL UNIQUE,
    password varchar(100) NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transaction (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    book_id uuid NULL,
    borrower_id uuid NULL,
    status varchar(50) NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    admin_id uuid NULL,
    CONSTRAINT fk_book FOREIGN KEY(book_id) REFERENCES book(id),
    CONSTRAINT fk_borrower FOREIGN KEY(borrower_id) REFERENCES borrower(id),
    CONSTRAINT fk_admin FOREIGN KEY(admin_Id) REFERENCES admin(id)
);

INSERT INTO book (title,author,release_year)
VALUES('Kenanganku','Maya','2011');

INSERT INTO book (title,author,release_year)
VALUES('Lomba 17','Agung','2008');

INSERT INTO book (title,author,release_year)
VALUES('Libra','Septi','2012');

INSERT INTO book (title,author,release_year)
VALUES('Peluk aku','Maya','2013');

INSERT INTO book (title,author,release_year)
VALUES('Ojek gila','Aji','2014');

INSERT INTO book (title,author,release_year)
VALUES('Barongsai mengaung','Jarwo','2017');

INSERT INTO admin (name,email,password)
VALUES('Ade','ade@mail.com','passwords');

INSERT INTO borrower (name,email,password, address)
VALUES('Topan','Topan@mail.com','passwords', 'Surabaya');