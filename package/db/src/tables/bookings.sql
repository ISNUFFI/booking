CREATE TABLE bookings (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    slot_id     BIGINT NOT NULL REFERENCES slots(id) ON DELETE CASCADE,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),

    UNIQUE(slot_id)
);
