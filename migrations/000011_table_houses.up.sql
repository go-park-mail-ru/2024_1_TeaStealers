DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'statusAreaHouse') THEN
        CREATE TYPE statusAreaHouse AS ENUM ('IHC', 'DNP', 'G', 'F', 'PSP');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'statusHomeHouse') THEN
        CREATE TYPE statusHomeHouse AS ENUM ('Live', 'RepairNeed', 'CompleteNeed', 'Renovation');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS houses (
    id UUID NOT NULL PRIMARY KEY,
    buildingId UUID NOT NULL REFERENCES buildings(id) ON DELETE CASCADE,
    advertTypeId UUID NOT NULL REFERENCES advertTypes(id) ON DELETE CASCADE,
    ceilingHeight FLOAT DEFAULT NULL,
    squareArea FLOAT DEFAULT NULL,
    squareHouse FLOAT DEFAULT NULL,
    bedroomCount INT DEFAULT NULL,
    statusArea statusAreaHouse DEFAULT NULL,
    cottage BOOLEAN DEFAULT NULL,
    statusHome statusHomeHouse DEFAULT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);