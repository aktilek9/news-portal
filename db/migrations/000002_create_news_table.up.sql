CREATE TABLE news(
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    content     TEXT NOT NULL,
    author_id   INTEGER NOT NULL REFERENCES users(id),
    created_at  TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at  TIMESTAMP NOT NULL DEFAULT current_timestamp,
    deleted_at  TIMESTAMP,

    CONSTRAINT fk_news_author
        FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_news_author_id ON news(author_id);
CREATE INDEX idx_news_deleted_at ON news(deleted_at);
