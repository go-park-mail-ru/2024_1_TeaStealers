CREATE TABLE IF NOT EXISTS user_data (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    password_hash TEXT CONSTRAINT password_hash_length CHECK ( char_length(password_hash) <= 40) NOT NULL,
    level_update INTEGER NOT NULL DEFAULT 1,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20 AND char_length(phone) >= 1) NOT NULL UNIQUE,
    email TEXT CONSTRAINT email_length CHECK ( char_length(email) <= 255 AND char_length(email) >= 1) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS advert (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user_data(id),
    title TEXT CONSTRAINT title_length CHECK ( char_length(title) <= 127) NOT NULL,
    type_placement TEXT CONSTRAINT type_placement_length CHECK ( char_length(type_placement) <= 6 AND type_placement IN ('Rent', 'Sale')) NOT NULL,
    description TEXT CONSTRAINT description_length CHECK ( char_length(description) <= 1500) NOT NULL,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20) NOT NULL,
    is_agent BOOLEAN NOT NULL DEFAULT FALSE,
    priority SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );


CREATE TABLE IF NOT EXISTS favourite_advert (
                                                user_id BIGINT NOT NULL REFERENCES user_data(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, advert_id)
    );