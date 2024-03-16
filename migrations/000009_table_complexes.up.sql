DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'classHouse') THEN
        CREATE TYPE classHouse AS ENUM ('Econom', 'Comfort', 'Business', 'Premium', 'Elite');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS complexes (
    id UUID PRIMARY KEY,
    companyId UUID NOT NULL REFERENCES companies(id),
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 255) NOT NULL UNIQUE,    
    adress TEXT CONSTRAINT adress_length CHECK ( char_length(adress) <= 512) NOT NULL,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    description TEXT NOT NULL,
    dateBeginBuild DATE NOT NULL,
    dateEndBuild DATE NOT NULL,
    withoutFinishingOption BOOLEAN DEFAULT NULL,
    finishingOption BOOLEAN DEFAULT NULL,
    preFinishingOption BOOLEAN DEFAULT NULL,
    classHousing classHouse DEFAULT NULL,
    parking BOOLEAN DEFAULT NULL,
    security BOOLEAN DEFAULT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);