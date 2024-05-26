CREATE TABLE merchants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    merchant_category VARCHAR(20) NOT NULL CHECK (merchant_category IN ('SmallRestaurant', 'MediumRestaurant', 'LargeRestaurant', 'MerchandiseRestaurant', 'BoothKiosk', 'ConvenienceStore')),
    image_url TEXT NOT NULL,
    location_lat FLOAT NOT NULL,
    location_long FLOAT NOT NULL
);
