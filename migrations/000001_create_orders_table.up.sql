CREATE TYPE status AS ENUM ('NEW', 'DELIVERED');

CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(15) PRIMARY KEY,
    quantity INT NOT NULL,
    title VARCHAR(15) NOT NULL,
    comment TEXT,
    uploaded_at timestamp,
    current_status status,
);