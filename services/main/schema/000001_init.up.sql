CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    first_name    varchar(255),
    last_name     varchar(255),
    birth_date    varchar(255),
    email         varchar(255),
    phone_number  varchar(255)
);
