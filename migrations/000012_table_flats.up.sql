CREATE TABLE IF NOT EXISTS flats (
    id UUID NOT NULL PRIMARY KEY,
    buildingId UUID NOT NULL REFERENCES buildings(id) ON DELETE CASCADE,
    advertTypeId UUID NOT NULL REFERENCES advertTypes(id) ON DELETE CASCADE,
    floor SMALLINT DEFAULT NULL,
    ceilingHeight FLOAT DEFAULT NULL,
    squareGeneral FLOAT DEFAULT NULL,
    roomCount SMALLINT DEFAULT NULL,
    squareResidential FLOAT DEFAULT NULL,
    apartament BOOLEAN DEFAULT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);