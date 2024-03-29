CREATE TABLE IF NOT EXISTS user_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    session_id VARCHAR(255) UNIQUE NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    user_agent TEXT,
    status boolean,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deaded_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (user_id, session_id)
);