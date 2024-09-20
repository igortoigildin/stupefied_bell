CREATE TABLE IF NOT EXISTS orders (
    number VARCHAR(15) PRIMARY KEY,
    quantity INT NOT NULL,
    title VARCHAR(15) NOT NULL,
    comment TEXT,
    uploaded_at timestamp
);