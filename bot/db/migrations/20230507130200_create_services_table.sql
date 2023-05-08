CREATE TABLE IF NOT EXISTS services
(
    id           SERIAL PRIMARY KEY,
    id_tg        INTEGER                  NOT NULL REFERENCES users (id_tg),
    service_name TEXT                     NOT NULL,
    login        TEXT                     NOT NULL,
    password     TEXT                     NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
