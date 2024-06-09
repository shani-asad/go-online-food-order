CREATE TABLE item_orders (
    item_id SERIAL,
    order_id INT NOT NULL,
    quantity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);