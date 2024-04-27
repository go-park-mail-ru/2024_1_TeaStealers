CREATE TABLE IF NOT EXISTS  question_answer (
    user_id UUID NOT NULL REFERENCES users(id),
    question_id UUID NOT NULL REFERENCES question(id),
    mark INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, question_id)
);