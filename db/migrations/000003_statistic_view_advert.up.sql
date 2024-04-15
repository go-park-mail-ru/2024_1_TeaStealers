CREATE TABLE IF NOT EXISTS statistic_view_advert (
    user_id BIGINT NOT NULL REFERENCES user_data(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, advert_id)
);