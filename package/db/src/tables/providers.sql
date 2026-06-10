CREATE TABLE providers (
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        TEXT NOT NULL,
    owner       INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    description TEXT,
    created_at  TIMESTAMP NOT NULL DEFAULT now()
);
