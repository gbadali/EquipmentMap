CREATE TABLE equipment (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    parent INTEGER,
    FOREIGN KEY (parent) REFERENCES equipment(id)
);