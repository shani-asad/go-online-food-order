CREATE TABLE orde (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    merchant_id INT NOT NULL,
    already_placed BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);