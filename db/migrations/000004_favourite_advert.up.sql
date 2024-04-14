CREATE TABLE IF NOT EXISTS favourite_advert (
    user_id BIGINT NOT NULL REFERENCES table_user(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE(user_id, advert_id)
);