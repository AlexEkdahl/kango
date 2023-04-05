-- init.sql

CREATE TABLE IF NOT EXISTS task (
    id BIGSERIAL PRIMARY KEY,
    status INTEGER NOT NULL,
    subject VARCHAR(255) NOT NULL,
    description TEXT
);
