CREATE SCHEMA subscription_service;

CREATE TABLE subscription_service.subscriptions (
    id              SERIAL              PRIMARY KEY,
    service_name    VARCHAR(100)        NOT NULL CHECK(char_length(service_name) BETWEEN 1 AND 100),
    price           INTEGER             NOT NULL CHECK(price > 0),
    user_id         UUID                NOT NULL,
    start_date      DATE                NOT NULL,
    end_date        DATE                NOT NULL,
    created_at      TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
)