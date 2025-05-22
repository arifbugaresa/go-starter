-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE tm_roles (
    role_id         SERIAL PRIMARY KEY,
    role_name       VARCHAR(50) NOT NULL UNIQUE,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      INT DEFAULT NULL,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INT DEFAULT NULL
);

-- Insert roles
INSERT INTO tm_roles (role_name)
VALUES
    ('admin'),
    ('user');

-- +migrate StatementEnd
