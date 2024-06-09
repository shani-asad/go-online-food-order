CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    already_placed BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);