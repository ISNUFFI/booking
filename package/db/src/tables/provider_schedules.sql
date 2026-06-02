CREATE TABLE provider_schedules (
    user_id     BIGINT      NOT NULL REFERENCES users(id),
    weekday     SMALLINT    NOT NULL,
    start_time  TIME        NOT NULL,
    end_time    TIME        NOT NULL,

    PRIMARY KEY (user_id, weekday)
);
