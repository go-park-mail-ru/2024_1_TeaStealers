CREATE TABLE IF NOT EXISTS companies (
    id UUID NOT NULL PRIMARY KEY,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    name TEXT CONSTRAINT name_length CHECK ( char_length(photo) <= 255) NOT NULL UNIQUE,
    yearFounded SMALLINT NOT NULL,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(photo) <= 20) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);