CREATE TABLE IF NOT EXISTS question (
    id UUID NOT NULL PRIMARY KEY,
    question_text TEXT CONSTRAINT question_text_length CHECK ( char_length(question_text) <= 120 AND char_length(question_text) >= 1) NOT NULL,
    theme TEXT CONSTRAINT theme_length CHECK ( char_length(theme) <= 15 AND theme IN ('mainPage', 'createAdvert', 'filterPage', 'profile', 'myAdverts')) NOT NULL,
    max_mark INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);
