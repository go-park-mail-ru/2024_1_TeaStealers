CREATE TABLE IF NOT EXISTS favourite_advert (
    user_id BIGINT NOT NULL REFERENCES user_data(id),
    advert_id BIGINT NOT NULL REFERENCES advert(id),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, advert_id)
);

CREATE OR REPLACE FUNCTION update_likes_on_insert()
RETURNS TRIGGER AS $$
BEGIN
UPDATE advert
SET likes = likes + 1
WHERE id = NEW.advert_id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_favourite_insert
    AFTER INSERT ON favourite_advert
    FOR EACH ROW
    EXECUTE FUNCTION update_likes_on_insert();





CREATE OR REPLACE FUNCTION update_likes_on_delete()
RETURNS TRIGGER AS $$
BEGIN
UPDATE advert
SET likes = likes - 1
WHERE id = OLD.advert_id;
RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_favourite_delete
    AFTER DELETE ON favourite_advert
    FOR EACH ROW
    EXECUTE FUNCTION update_likes_on_delete();




CREATE OR REPLACE FUNCTION update_likes_on_advert_delete()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.is_deleted = TRUE THEN
UPDATE advert
SET likes = likes - 1
WHERE id = OLD.id;
END IF;
RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_advert_delete
    AFTER UPDATE ON advert
    FOR EACH ROW
    EXECUTE FUNCTION update_likes_on_advert_delete();