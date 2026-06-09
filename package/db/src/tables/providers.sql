CREATE TABLE providers (
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT,
    created_at  TIMESTAMP NOT NULL DEFAULT now()
);
