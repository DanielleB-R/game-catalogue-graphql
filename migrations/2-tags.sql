CREATE TABLE IF NOT EXISTS tag (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS game_tag (
       tag_id INTEGER REFERENCES tag NOT NULL,
       game_id INTEGER REFERENCES game NOT NULL,
       PRIMARY KEY (tag_id, game_id)
);
