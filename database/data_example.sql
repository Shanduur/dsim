-- insertion of examplary user:

INSERT INTO users VALUES(
    DEFAULT,
    'admin',
    'admin@example.org',
    '1999-12-31'
);

-- insertion of examplary type:

INSERT INTO types VALUES(
    DEFAULT,
    'txt',
    'Standard text document that contains unformatted text.'
);

-- insertion of examplary blob:

INSERT INTO blobs VALUES(
    DEFAULT,
    DECODE('706C75676162626C20697320617765736F6D650A0A', 'hex'),
    1,
    'example',
    1,
    '2020-08-10'
);

-- list database users:

SELECT * FROM pg_catalog.pg_user;

-- display some info to check integrity:

SELECT u.user_id, u.username, u.email, b.blob_name, t.type_extension, encode(b.blob_data, 'hex') AS blob_data
FROM users AS u
JOIN blobs AS b ON u.user_id=b.owner_id
JOIN types AS t ON t.type_id=b.blob_type;
