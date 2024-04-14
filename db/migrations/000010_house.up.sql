DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'statusArea') THEN
        CREATE TYPE statusArea AS ENUM ('IHC', 'DNP', 'G', 'F', 'PSP', 'None');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'statusHomeHouse') THEN
        CREATE TYPE statusHomeHouse AS ENUM ('Live', 'RepairNeed', 'CompleteNeed', 'Renovation', 'None');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS house (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    building_id BIGINT NOT NULL REFERENCES building(id),
    ceiling_height FLOAT NOT NULL,
    square_area FLOAT NOT NULL,
    square_house FLOAT NOT NULL,
    bedroom_count INT NOT NULL,
    status_area_house statusArea NOT NULL DEFAULT 'None',
    cottage BOOLEAN NOT NULL DEFAULT FALSE,
    status_home_house statusHomeHouse NOT NULL DEFAULT 'None',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);