CREATE TABLE comment(
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER NOT NULL REFERENCES users(id),
    news_id INTEGER NOT NULL REFERENCES news(id),
    created_at  TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at  TIMESTAMP NOT NULL DEFAULT current_timestamp,
    deleted_at  TIMESTAMP NULL,

    CONSTRAINT fk_comment_author
        FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_comment_news
        FOREIGN KEY (news_id)
        REFERENCES news(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_comment_author ON comment(author_id);
CREATE INDEX idx_comment_news ON comment(news_id);