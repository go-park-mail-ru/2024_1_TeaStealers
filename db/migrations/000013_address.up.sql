CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS address (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    address_point GEOGRAPHY(Point, 4326) NOT NULL UNIQUE,
    house_name_id BIGINT NOT NULL REFERENCES house_name(id),
    metro TEXT CONSTRAINT metro_length CHECK ( char_length(metro) <= 120) NOT NULL
);

INSERT INTO province (name) VALUES ('California');
INSERT INTO province (name) VALUES ('New York');

-- Inserting into town table
INSERT INTO town (province_id, name) VALUES (1, 'Los Angeles');
INSERT INTO town (province_id, name) VALUES (1, 'San Francisco');
INSERT INTO town (province_id, name) VALUES (2, 'New York City');
INSERT INTO town (province_id, name) VALUES (2, 'Buffalo');

-- Inserting into street table
INSERT INTO street (town_id, name) VALUES (1, 'Sunset Boulevard');
INSERT INTO street (town_id, name) VALUES (1, 'Hollywood Boulevard');
INSERT INTO street (town_id, name) VALUES (2, 'Lombard Street');
INSERT INTO street (town_id, name) VALUES (3, 'Broadway');
INSERT INTO street (town_id, name) VALUES (4, 'Main Street');

-- Inserting into house_name table
INSERT INTO house_name (street_id, name) VALUES (1, '1234');
INSERT INTO house_name (street_id, name) VALUES (1, '5678');
INSERT INTO house_name (street_id, name) VALUES (2, '4321');
INSERT INTO house_name (street_id, name) VALUES (3, '9876');
INSERT INTO house_name (street_id, name) VALUES (4, '6543');

-- Inserting into address table with a test value for address_point
INSERT INTO address (address_point, house_name_id, metro) VALUES (ST_GeographyFromText('POINT(-118.2437 34.0522)'), 1, 'Metro Station A');
INSERT INTO address (address_point, house_name_id, metro) VALUES (ST_GeographyFromText('POINT(-122.4194 37.7749)'), 2, 'Metro Station B');
INSERT INTO address (address_point, house_name_id, metro) VALUES (ST_GeographyFromText('POINT(-122.4313 37.8044)'), 3, 'Metro Station C');
INSERT INTO address (address_point, house_name_id, metro) VALUES (ST_GeographyFromText('POINT(-73.9352 40.7306)'), 4, 'Metro Station D');
INSERT INTO address (address_point, house_name_id, metro) VALUES (ST_GeographyFromText('POINT(-78.8784 42.8864)'), 5, 'Metro Station E');