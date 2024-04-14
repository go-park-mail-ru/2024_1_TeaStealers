CREATE TABLE IF NOT EXISTS image (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    priority SMALLINT NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE(advert_id, priority)
);