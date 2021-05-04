-- Database definition for PostgreSQL

CREATE DATABASE dsime_db;

CREATE TABLE users (
    user_id BIGSERIAL NOT NULL,
    user_name VARCHAR(32) NOT NULL,
    user_key BYTEA NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    PRIMARY KEY(user_id)
);

CREATE TABLE filetypes (
    type_id SERIAL NOT NULL,
    type_extension VARCHAR(16) NOT NULL,
    type_description VARCHAR(1024) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    PRIMARY KEY(type_id)
);

CREATE TABLE blobs (
    blob_id BIGSERIAL NOT NULL,
    blob_data BYTEA NOT NULL,
    blob_type INT REFERENCES filetypes(type_id),
    blob_name VARCHAR(256),
    blob_checksum BYTEA NOT NULL,
    insertion_date DATE NOT NULL,
    parents INT[],
    active BOOLEAN DEFAULT TRUE,
    PRIMARY KEY(blob_id)
);

CREATE TABLE nodes (
    node_ip VARCHAR(15) NOT NULL,
    node_port INT NOT NULL,
    node_reg_jobs INT,
    node_max_jobs INT NOT NULL,
    node_timeout TIMESTAMP NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    PRIMARY KEY(node_ip, node_port, active)
);
