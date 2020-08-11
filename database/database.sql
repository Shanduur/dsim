-- Database definition for PostgreSQL

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(32) NOT NULL,
    email VARCHAR(64),
    creation_date DATE NOT NULL
);

CREATE TABLE types (
    type_id SERIAL PRIMARY KEY NOT NULL,
    type_extension VARCHAR(16) NOT NULL,
    type_description VARCHAR(1024) NOT NULL
);

CREATE TABLE blobs (
    blob_id BIGSERIAL PRIMARY KEY NOT NULL,
    blob_data BYTEA NOT NULL,
    blob_type INT REFERENCES types(type_id),
    blob_name VARCHAR(256),
    owner_id BIGINT REFERENCES users(user_id) NOT NULL,
    insertion_date DATE NOT NULL
);
