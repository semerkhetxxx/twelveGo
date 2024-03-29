CREATE TABLE IF NOT EXISTS images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    extension VARCHAR(10) NOT NULL,
    image_data BYTEA,
    path VARCHAR(255),
    tags JSONB
);
