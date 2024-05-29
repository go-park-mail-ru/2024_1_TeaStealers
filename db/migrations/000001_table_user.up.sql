CREATE TABLE IF NOT EXISTS user_data (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    password_hash TEXT CONSTRAINT password_hash_length CHECK ( char_length(password_hash) <= 40) NOT NULL,
    level_update INTEGER NOT NULL DEFAULT 1,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20 AND char_length(phone) >= 1) NOT NULL UNIQUE,
    email TEXT CONSTRAINT email_length CHECK ( char_length(email) <= 255 AND char_length(email) >= 1) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);