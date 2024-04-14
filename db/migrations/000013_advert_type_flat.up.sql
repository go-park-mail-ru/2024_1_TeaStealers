CREATE TABLE IF NOT EXISTS advert_type_flat (
    flat_id BIGINT NOT NULL REFERENCES flat(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (flat_id, advert_id)
);