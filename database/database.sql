-- Database definition for PostgreSQL

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY NOT NULL,
    user_name VARCHAR(32) NOT NULL,
    user_key BYTEA NOT NULL
);

CREATE TABLE filetypes (
    type_id SERIAL PRIMARY KEY NOT NULL,
    type_extension VARCHAR(16) NOT NULL,
    type_description VARCHAR(1024) NOT NULL
);

CREATE TABLE blobs (
    blob_id BIGSERIAL PRIMARY KEY NOT NULL,
    blob_data BYTEA NOT NULL,
    blob_type INT REFERENCES filetypes(type_id),
    blob_name VARCHAR(256),
    owner_id BIGINT REFERENCES users(user_id) NOT NULL,
    insertion_date DATE NOT NULL
);

CREATE TABLE nodes (
    node_ip VARCHAR(15) PRIMARY KEY NOT NULL,
    node_port INT NOT NULL,
    node_reg_jobs INT,
    node_max_jobs INT NOT NULL,
    node_timeout TIMESTAMP NOT NULL
)
