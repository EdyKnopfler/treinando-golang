-- +goose Up

CREATE TABLE banhos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    id_pet INTEGER NOT NULL,
    data TEXT NOT NULL,
    valor REAL NOT NULL,
    FOREIGN KEY (id_pet) REFERENCES pets(id)
);

-- +goose Down

DROP TABLE banhos;