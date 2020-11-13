CREATE TABLE IF NOT EXISTS platform (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS game (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL,
       platform INTEGER REFERENCES platform NOT NULL
);
