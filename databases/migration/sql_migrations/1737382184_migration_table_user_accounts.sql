-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE tm_accounts (
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(100) NOT NULL UNIQUE,
    email           VARCHAR(100) NOT NULL UNIQUE,
    password        VARCHAR(255) NOT NULL,
    full_name       VARCHAR(100),
    photo           TEXT,
    role_id         INT NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      INT DEFAULT NULL,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INT DEFAULT NULL,
    FOREIGN KEY (role_id) REFERENCES tm_roles(role_id)
);

-- +migrate StatementEnd