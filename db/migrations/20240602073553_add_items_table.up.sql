CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    product_category VARCHAR(20) NOT NULL CHECK (product_category IN ('Beverage', 'Food', 'Snack', 'Condiments', 'Additions')),
    price INT NOT NULL,
    image_url TEXT NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp
);
