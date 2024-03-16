DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'typePlacementAdvert') THEN
        CREATE TYPE typePlacementAdvert AS ENUM ('House', 'Flat');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS adverts (
    id UUID NOT NULL PRIMARY KEY,
    userId UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    advertTypeId UUID NOT NULL REFERENCES advertTypes(id) ON DELETE CASCADE,
    advertTypePlacement typePlacementAdvert NOT NULL,
    title TEXT CONSTRAINT title_length CHECK ( char_length(title) <= 127) NOT NULL,
    description TEXT NOT NULL,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20) NOT NULL UNIQUE,
    isAgent BOOLEAN NOT NULL,
    priority SMALLINT NOT NULL DEFAULT 1,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE
);
