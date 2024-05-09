CREATE TABLE IF NOT EXISTS user_profile_data (
    user_id BIGINT PRIMARY KEY NOT NULL REFERENCES user_data(id),
    first_name TEXT CONSTRAINT first_name_length CHECK ( char_length(first_name) <= 127) NOT NULL,
    surname TEXT CONSTRAINT surname_length CHECK ( char_length(surname) <= 127) NOT NULL,
    birthdate DATE DEFAULT NULL,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);