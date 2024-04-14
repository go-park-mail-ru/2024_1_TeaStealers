CREATE TABLE IF NOT EXISTS advert (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES table_user(id),
    title TEXT CONSTRAINT title_length CHECK ( char_length(title) <= 127) NOT NULL DEFAULT '',
    description TEXT CONSTRAINT description_length CHECK ( char_length(description) <= 1500) NOT NULL DEFAULT '',
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20) NOT NULL DEFAULT '',
    is_agent BOOLEAN NOT NULL DEFAULT FALSE,
    priority SMALLINT NOT NULL DEFAULT 1,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);