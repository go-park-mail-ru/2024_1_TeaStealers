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
    id UUID NOT NULL PRIMARY KEY
);