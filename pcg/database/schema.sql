CREATE TABLE IF NOT EXISTS news (
    id SERIAL PRIMARY KEY,
    title TEXT UNIQUE,
    description TEXT,
    INTEGER DEFAULT 0,
    source TEXT
);