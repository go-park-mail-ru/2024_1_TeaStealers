CREATE TABLE IF NOT EXISTS complex (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_id BIGINT NOT NULL REFERENCES company(id),
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 255) NOT NULL UNIQUE,    
    address TEXT CONSTRAINT address_length CHECK ( char_length(address) <= 512) NOT NULL,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    description TEXT CONSTRAINT description_length CHECK ( char_length(description) <= 1500) NOT NULL,
    date_begin_build DATE NOT NULL,
    date_end_build DATE NOT NULL,
    without_finishing_option BOOLEAN NOT NULL DEFAULT FALSE,
    finishing_option BOOLEAN NOT NULL DEFAULT FALSE,
    pre_finishing_option BOOLEAN NOT NULL DEFAULT FALSE,
    class_housing TEXT CONSTRAINT class_housing_length CHECK ( char_length(class_housing) <= 10 AND class_housing IN ('Econom', 'Comfort', 'Business', 'Premium', 'Elite', 'None')) NOT NULL,
    parking BOOLEAN NOT NULL DEFAULT FALSE,
    security BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);