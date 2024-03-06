CREATE TABLE IF NOT EXISTS buildings (
    id             UUID PRIMARY KEY,
    location          VARCHAR(250) NOT NULL,
    description  TEXT NOT NULL,
    data_creation DATE NOT NULL,
    is_deleted BOOLEAN NOT NULL
);