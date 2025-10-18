CREATE TABLE IF NOT EXISTS employees (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    position varchar(255) NOT NULL,
    salary double precision NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);