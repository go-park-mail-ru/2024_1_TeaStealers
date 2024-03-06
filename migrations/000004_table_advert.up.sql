CREATE TABLE IF NOT EXISTS adverts (
    id             UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) NOT NULL,
    phone INTEGER NOT NULL,
    description TEXT NOT NULL,
    building_id UUID NULL,
    company_id UUID  NULL,
    price FLOAT NOT NULL,
    location VARCHAR(350) NOT NULL,
    data_creation DATE NOT NULL,
    is_deleted BOOLEAN NOT NULL
);