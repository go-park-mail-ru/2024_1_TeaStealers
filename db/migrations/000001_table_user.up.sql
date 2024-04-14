CREATE TABLE IF NOT EXISTS "user" (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    password_hash TEXT CONSTRAINT password_hash_length CHECK ( char_length(password_hash) <= 40) NOT NULL,
    level_update INTEGER NOT NULL DEFAULT 1,
    first_name TEXT CONSTRAINT first_name_length CHECK ( char_length(first_name) <= 127) NOT NULL,
    second_name TEXT CONSTRAINT second_name_length CHECK ( char_length(second_name) <= 127) NOT NULL,
    date_birthday DATE DEFAULT NULL,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20 AND char_length(phone) >= 1) NOT NULL UNIQUE,
    email TEXT CONSTRAINT email_length CHECK ( char_length(email) <= 255 AND char_length(email) >= 1) NOT NULL UNIQUE,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);