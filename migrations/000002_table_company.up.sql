CREATE TABLE IF NOT EXISTS companies (
    id             UUID PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    phone INTEGER NOT NULL,
    description TEXT NOT NULL,
    data_creation DATE NOT NULL,
    is_deleted BOOLEAN NOT NULL
);