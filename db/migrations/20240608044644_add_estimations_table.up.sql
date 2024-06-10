CREATE TABLE estimations (
    id SERIAL PRIMARY KEY,
    order_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX fdm ON estimations (id);