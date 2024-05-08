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