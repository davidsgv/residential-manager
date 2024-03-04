/*
    script: create schema
    author: DavidSG
    description: initial schema for database
*/

CREATE TABLE permit(
    id INT PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
);

CREATE TABLE rol(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE permit_rol(
    permit_id INT NOT NULL,
    rol_id INT NOT NULL,
    PRIMARY KEY (permit_id, rol_id)
);

CREATE TABLE block(
    id SERIAL PRIMARY KEY,
    block VARCHAR NOT NULL UNIQUE
);

CREATE TABLE apartment(
    id SERIAL PRIMARY KEY,
    number VARCHAR NOT NULL,
    block_id INT NOT NULL,
    CONSTRAINT fk_aparment_block
        FOREIGN KEY(block_id) 
            REFERENCES block(id),
    UNIQUE(number, block_id)
);

CREATE TABLE user_account(
    id SERIAL PRIMARY KEY,
    uuid VARCHAR NOT NULL UNIQUE,
    mail VARCHAR NOT NULL UNIQUE,
    password VARCHAR,
    verified BOOLEAN NOT NULL,
    token VARCHAR NOT NULL UNIQUE,
    token_expire TIMESTAMP NOT NULL,
    apartment_id INT NOT NULL,
    rol_id INT NOT NULL,
    CONSTRAINT fk_user_account_apartment
        FOREIGN KEY(apartment_id) 
            REFERENCES apartment(id),
    CONSTRAINT fk_user_account_rol
        FOREIGN KEY(rol_id) 
            REFERENCES rol(id)
);