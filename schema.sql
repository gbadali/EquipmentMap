CREATE TABLE equipment (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    parent INTEGER NOT NULL,
    FOREIGN KEY (parent) REFERENCES equipment(id)
);

-- TODO: add a user table