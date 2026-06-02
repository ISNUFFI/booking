CREATE TYPE user_role AS ENUM (
    'user',
    'provider'
);

CREATE TABLE users (
    id              BIGSERIAL                   PRIMARY KEY,
    email           TEXT                        NOT NULL UNIQUE,
    role            user_role                   NOT NULL DEFAULT 'user',
    password_hash   TEXT                        NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT now()
);
