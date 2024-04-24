CREATE TABLE IF NOT EXISTS advert_type_house (
    house_id BIGINT NOT NULL REFERENCES house(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (house_id, advert_id)
);