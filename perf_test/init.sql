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

CREATE TABLE IF NOT EXISTS statistic_view_advert (
                                                     user_id BIGINT NOT NULL REFERENCES user_data(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, advert_id)
    );

CREATE TABLE IF NOT EXISTS favourite_advert (
                                                user_id BIGINT NOT NULL REFERENCES user_data(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, advert_id)
    );

CREATE TABLE IF NOT EXISTS image (
                                     id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                     advert_id BIGINT NOT NULL REFERENCES advert(id),
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    priority SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE(advert_id, priority)
    );

CREATE TABLE IF NOT EXISTS price_change (
                                            id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                            advert_id BIGINT NOT NULL REFERENCES advert(id),
    price BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS company (
                                       id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                       photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 255) NOT NULL UNIQUE,
    creation_year SMALLINT NOT NULL,
    phone TEXT CONSTRAINT phone_length CHECK ( char_length(phone) <= 20) NOT NULL UNIQUE,
    description TEXT CONSTRAINT description_length CHECK ( char_length(description) <= 1500) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

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

CREATE TABLE IF NOT EXISTS province (
                                        id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                        name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 120) NOT NULL
    );

CREATE TABLE IF NOT EXISTS town (
                                    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                    province_id BIGINT NOT NULL REFERENCES province(id),
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 60) NOT NULL
    );

CREATE TABLE IF NOT EXISTS street (
                                      id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                      town_id BIGINT NOT NULL REFERENCES town(id),
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 120) NOT NULL
    );

CREATE TABLE IF NOT EXISTS house_name (
                                          id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                          street_id BIGINT NOT NULL REFERENCES street(id),
    name TEXT CONSTRAINT name_length CHECK ( char_length(name) <= 40) NOT NULL
    );

CREATE TABLE IF NOT EXISTS address (
                                       id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                       address_point GEOGRAPHY(Point, 4326) NOT NULL UNIQUE,
    house_name_id BIGINT NOT NULL REFERENCES house_name(id),
    metro TEXT CONSTRAINT metro_length CHECK ( char_length(metro) <= 120) NOT NULL
    );

CREATE TABLE IF NOT EXISTS building (
                                        id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                        complex_id BIGINT NULL REFERENCES complex(id),
    address_id BIGINT NULL REFERENCES address(id),
    floor SMALLINT NOT NULL,
    material_building TEXT CONSTRAINT material_building_length CHECK ( char_length(material_building) <= 25 AND material_building IN ('Brick', 'Monolithic', 'Wood', 'Panel', 'Stalinsky', 'Block', 'MonolithicBlock', 'Frame', 'AeratedConcreteBlock', 'GasSilicateBlock', 'FoamСoncreteBlock', 'None')) NOT NULL,
    year_creation SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS house (
                                     id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                     building_id BIGINT NOT NULL REFERENCES building(id),
    ceiling_height FLOAT NOT NULL,
    square_area FLOAT NOT NULL,
    square_house FLOAT NOT NULL,
    bedroom_count INT NOT NULL,
    status_area_house TEXT CONSTRAINT status_area_house_length CHECK ( char_length(status_area_house) <= 5 AND status_area_house IN ('IHC', 'DNP', 'G', 'F', 'PSP', 'None')) NOT NULL,
    cottage BOOLEAN NOT NULL DEFAULT FALSE,
    status_home_house TEXT CONSTRAINT status_home_house_length CHECK ( char_length(status_home_house) <= 15 AND status_home_house IN ('Live', 'RepairNeed', 'CompleteNeed', 'Renovation', 'None')) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS flat (
                                    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                    building_id BIGINT NOT NULL REFERENCES building(id),
    floor SMALLINT NOT NULL,
    ceiling_height FLOAT NOT NULL,
    square_general FLOAT NOT NULL,
    bedroom_count SMALLINT NOT NULL,
    square_residential FLOAT NOT NULL,
    apartament BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS advert_type_house (
                                                 house_id BIGINT NOT NULL REFERENCES house(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (house_id, advert_id)
    );

CREATE TABLE IF NOT EXISTS advert_type_flat (
                                                flat_id BIGINT NOT NULL REFERENCES flat(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (flat_id, advert_id)
    );

CREATE TABLE IF NOT EXISTS question (
                                        id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                        question_text TEXT CONSTRAINT question_text_length CHECK ( char_length(question_text) <= 120 AND char_length(question_text) >= 1) NOT NULL,
    theme TEXT CONSTRAINT theme_length CHECK ( char_length(theme) <= 15 AND theme IN ('mainPage', 'createAdvert', 'filterPage', 'profile', 'myAdverts')) NOT NULL,
    max_mark INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS  question_answer (
                                                user_id BIGINT NOT NULL REFERENCES user_data(id),
    question_id BIGINT NOT NULL REFERENCES question(id),
    mark INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, question_id)
    );

CREATE TABLE IF NOT EXISTS user_profile_data (
                                                 user_id BIGINT PRIMARY KEY NOT NULL REFERENCES user_data(id),
    first_name TEXT CONSTRAINT first_name_length CHECK ( char_length(first_name) <= 127) NOT NULL,
    surname TEXT CONSTRAINT surname_length CHECK ( char_length(surname) <= 127) NOT NULL,
    birthdate DATE DEFAULT NULL,
    photo TEXT CONSTRAINT photo_length CHECK ( char_length(photo) <= 255) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );