CREATE EXTENSION IF NOT EXISTS postgis;

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'materialBuilding') THEN
        CREATE TYPE materialBuilding AS ENUM ('Brick', 'Monolithic', 'Wood', 'Panel', 'Stalinsky', 'Block', 'MonolithicBlock', 'Frame', 'AeratedConcreteBlock', 'GasSilicateBlock', 'FoamСoncreteBlock');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS buildings (
    id UUID NOT NULL PRIMARY KEY,
    complexId UUID NULL REFERENCES complexes(id) ON DELETE SET NULL,
    floor SMALLINT NOT NULL,
    material materialBuilding DEFAULT NULL,
    adress TEXT CONSTRAINT adress_length CHECK ( char_length(adress) <= 255) NOT NULL UNIQUE,
    adressPoint GEOGRAPHY(Point, 4326) NOT NULL UNIQUE,
    yearCreation SMALLINT NOT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);