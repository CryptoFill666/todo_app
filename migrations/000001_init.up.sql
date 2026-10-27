CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users (
    id           SERIAL                PRIMARY KEY,
    version      INT          NOT NULL DEFAULT 1,
    full_name    VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    phone_number VARCHAR(13)           CHECK(
        phone_number ~ '^\+380[0-9]{9}$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    )
);

CREATE TABLE todoapp.task (
    id          SERIAL                PRIMARY KEY,
    version     INT          NOT NULL DEFAULT 1,
    user_id     INT          NOT NULL REFERENCES todoapp.users(id) ON DELETE CASCADE,
    title       VARCHAR(100) NOT NULL CHECK (char_length(title) BETWEEN 3 AND 100),
    description TEXT                  CHECK (char_length(description) BETWEEN 10 AND 1000),
    status      BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    done_at     TIMESTAMP
);