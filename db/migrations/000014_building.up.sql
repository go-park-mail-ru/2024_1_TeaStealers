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