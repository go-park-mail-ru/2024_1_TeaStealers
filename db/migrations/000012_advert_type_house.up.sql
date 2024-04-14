CREATE TABLE IF NOT EXISTS advert_type_house (
    house_id BIGINT NOT NULL REFERENCES house(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE(house_id, advert_id)
);