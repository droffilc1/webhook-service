CREATE TABLE IF NOT EXISTS events (
    id TEXT PRIMARY KEY,
    type TEXT NOT NULL,
    payload JSONB,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS endpoints (
    id TEXT PRIMARY KEY,
    url TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TYPE DELIVERY_STATUS AS ENUM (
    'success',
    'failed'
);

CREATE TABLE IF NOT EXISTS deliveries (
    id TEXT PRIMARY KEY,
    event_id TEXT NOT NULL,
    endpoint_id TEXT NOT NULL,
    status DELIVERY_STATUS NOT NULL,
    status_code INT,
    attempt INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_event
    FOREIGN KEY (event_id)
    REFERENCES events (id),
    CONSTRAINT fk_endpoint
    FOREIGN KEY (endpoint_id)
    REFERENCES endpoints (id)
);
