DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'classHouse') THEN
        CREATE TYPE classHouse AS ENUM ('Econom', 'Comfort', 'Business', 'Premium', 'Elite', 'None');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS complex (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_id BIGINT NOT NULL REFERENCES company(id),
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 255) NOT NULL UNIQUE,    
    address TEXT CONSTRAINT address_length CHECK ( char_length(address) <= 512) NOT NULL DEFAULT '',
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL DEFAULT '',
    description TEXT CONSTRAINT description_length CHECK ( char_length(description) <= 1500) NOT NULL DEFAULT '',
    date_begin_build DATE NOT NULL,
    date_end_build DATE NOT NULL,
    without_finishing_option BOOLEAN NOT NULL DEFAULT FALSE,
    finishing_option BOOLEAN NOT NULL DEFAULT FALSE,
    pre_finishing_option BOOLEAN NOT NULL DEFAULT FALSE,
    class_housing classHouse NOT NULL DEFAULT 'None',
    parking BOOLEAN NOT NULL DEFAULT FALSE,
    security BOOLEAN NOT NULL DEFAULT FALSE,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);