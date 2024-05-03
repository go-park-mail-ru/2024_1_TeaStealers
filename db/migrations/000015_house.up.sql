CREATE TABLE IF NOT EXISTS house (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    building_id BIGINT NOT NULL REFERENCES building(id),
    ceiling_height FLOAT NOT NULL,
    square_area FLOAT NOT NULL,
    square_house FLOAT NOT NULL,
    bedroom_count INT NOT NULL,
    status_area_house TEXT CONSTRAINT status_area_house_length CHECK ( char_length(status_area_house) <= 5 AND status_area_house IN ('IHC', 'DNP', 'G', 'F', 'PSP', 'None')) NOT NULL,
    cottage BOOLEAN NOT NULL DEFAULT FALSE,
    status_home_house TEXT CONSTRAINT status_home_house_length CHECK ( char_length(status_home_house) <= 15 AND status_home_house IN ('Live', 'RepairNeed', 'CompleteNeed', 'Renovation', 'None')) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);