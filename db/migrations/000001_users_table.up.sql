CREATE TABLE users (
    id int NOT NULL ,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    phone_number text NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL,

    PRIMARY KEY(id)
);
