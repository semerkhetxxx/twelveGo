CREATE TABLE IF NOT EXISTS urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    key TEXT UNIQUE,
    secret_key TEXT UNIQUE,
    target_url TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    clicks INT DEFAULT 0
);
