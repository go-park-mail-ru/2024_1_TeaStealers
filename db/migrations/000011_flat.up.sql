CREATE TABLE IF NOT EXISTS flat (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    building_id BIGINT NOT NULL REFERENCES building(id),
    floor SMALLINT NOT NULL,
    ceiling_height FLOAT NOT NULL,
    square_general FLOAT NOT NULL,
    room_count SMALLINT NOT NULL,
    square_residential FLOAT NOT NULL,
    apartament BOOLEAN NOT NULL DEFAULT FALSE,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);