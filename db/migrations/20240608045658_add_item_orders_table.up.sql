CREATE TABLE item_orders (
    item_id SERIAL,
    order_id INT NOT NULL,
    quantity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ds ON item_orders (item_id);
CREATE INDEX ghggg ON item_orders (order_id);
CREATE INDEX ssdfsd ON item_orders (quantity);