CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL PRIMARY KEY,
    passwordHash TEXT CONSTRAINT passwordHash_length CHECK ( char_length(passwordHash) <= 40) NOT NULL,
    levelUpdate INTEGER NOT NULL DEFAULT 1,
    firstName TEXT CONSTRAINT firstName_length CHECK ( char_length(firstName) <= 127) DEFAULT NULL,
    secondName TEXT CONSTRAINT secondName_length CHECK ( char_length(secondName) <= 127) DEFAULT NULL,
    dateBirthday DATE DEFAULT NULL,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20 AND char_length(phone) >= 1) NOT NULL UNIQUE,
    email TEXT CONSTRAINT email_length CHECK ( char_length(email) <= 255 AND char_length(email) >= 1) NOT NULL UNIQUE,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255),
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);