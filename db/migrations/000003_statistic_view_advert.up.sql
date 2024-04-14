CREATE TABLE IF NOT EXISTS statistic_view_advert (
    user_id BIGINT NOT NULL REFERENCES "user"(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, advert_id)
);