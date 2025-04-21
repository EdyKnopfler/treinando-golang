-- +goose Up

CREATE TABLE pets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    idade INTEGER NOT NULL
);

-- +goose Down

DROP TABLE pets;