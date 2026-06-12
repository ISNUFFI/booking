CREATE TABLE bookings (
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id     INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    slot_id     INT NOT NULL REFERENCES slots(id) ON DELETE CASCADE,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),

    UNIQUE(slot_id)
);
