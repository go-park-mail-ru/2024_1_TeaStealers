CREATE TABLE IF NOT EXISTS province (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 120) NOT NULL
);