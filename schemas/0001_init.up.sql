CREATE TABLE IF NOT EXISTS links (
    id serial PRIMARY KEY,
    original VARCHAR(255),
    short VARCHAR(50)
);