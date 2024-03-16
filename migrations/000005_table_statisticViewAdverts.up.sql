CREATE TABLE IF NOT EXISTS statisticViewAdverts (
    id UUID NOT NULL PRIMARY KEY,
    userId UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    advertId UUID NOT NULL REFERENCES adverts(id) ON DELETE CASCADE,
    dateCreation TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(userId, advertId)
);