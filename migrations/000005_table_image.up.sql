CREATE TABLE IF NOT EXISTS images (
    id             UUID PRIMARY KEY,
    advert_id UUID REFERENCES adverts(id),
    filename     VARCHAR(350) NOT NULL,
    priority INTEGER NOT NULL,
    data_creation DATE NOT NULL,
    is_deleted BOOLEAN NOT NULL
);