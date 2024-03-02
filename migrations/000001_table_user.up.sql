CREATE TABLE IF NOT EXISTS users (
    id             UUID PRIMARY KEY,
    login          VARCHAR(50) NOT NULL,
    password_hash  VARCHAR(255) NOT NULL
);