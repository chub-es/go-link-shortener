CREATE TABLE IF NOT EXISTS links(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    original_url VARCHAR(255) NOT NULL,
    short_url VARCHAR(255) NOT NULL UNIQUE DEFAULT generateshorturl(6)
);