CREATE TABLE IF NOT EXISTS images (
    id UUID NOT NULL PRIMARY KEY
    /*advertId UUID NOT NULL REFERENCES adverts(id) ON DELETE SET NULL,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    priority SMALLINT NOT NULL DEFAULT 1,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE(advertId, priority)*/
);