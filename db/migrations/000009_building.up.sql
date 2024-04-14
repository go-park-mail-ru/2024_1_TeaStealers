CREATE EXTENSION IF NOT EXISTS postgis;

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'material') THEN
        CREATE TYPE material AS ENUM ('Brick', 'Monolithic', 'Wood', 'Panel', 'Stalinsky', 'Block', 'MonolithicBlock', 'Frame', 'AeratedConcreteBlock', 'GasSilicateBlock', 'FoamСoncreteBlock', 'None');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS building (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    complex_id BIGINT NULL REFERENCES complex(id),
    floor SMALLINT NOT NULL,
    material_building material NOT NULL DEFAULT 'None',
    address TEXT CONSTRAINT address_length CHECK ( char_length(address) <= 255) NOT NULL UNIQUE,
    address_point GEOGRAPHY(Point, 4326) NOT NULL UNIQUE,
    year_creation SMALLINT NOT NULL,
    date_creation TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);