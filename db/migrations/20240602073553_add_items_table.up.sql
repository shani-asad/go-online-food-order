CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    merchant_id INTEGER,
    product_category VARCHAR(20) NOT NULL CHECK (product_category IN ('Beverage', 'Food', 'Snack', 'Condiments', 'Additions')),
    price INT NOT NULL,
    image_url TEXT NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp
);

-- Index on merchant_id to speed up joins with merchants table
CREATE INDEX idx_items_merchant_id ON items (merchant_id);

-- Index on name to speed up searches by item name
CREATE INDEX idx_items_name ON items (name);

-- Optionally, if you often search by product_category as well
CREATE INDEX idx_items_product_category ON items (product_category);