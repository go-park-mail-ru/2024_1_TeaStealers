CREATE TABLE IF NOT EXISTS company (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 255) NOT NULL UNIQUE,
    creation_year SMALLINT NOT NULL,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20) NOT NULL UNIQUE,
    description TEXT CONSTRAINT description_length CHECK ( char_length(description) <= 1500) NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);