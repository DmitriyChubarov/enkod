CREATE TABLE IF NOT EXISTS person (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    phone TEXT,
    first_name TEXT,
    last_name TEXT
);
