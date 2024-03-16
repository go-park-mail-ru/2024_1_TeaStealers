DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'advertTypeAdvert') THEN
        CREATE TYPE advertTypeAdvert AS ENUM ('house', 'flat');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS advertTypes (
    id UUID NOT NULL PRIMARY KEY,
    advertType advertTypeAdvert NOT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);