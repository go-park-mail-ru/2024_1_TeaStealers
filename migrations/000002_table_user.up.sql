CREATE TABLE IF NOT EXISTS users (
                                     id             UUID PRIMARY KEY,
                                     login          VARCHAR(50) NOT NULL UNIQUE ,
                                     password_hash  VARCHAR(255) NOT NULL
);