DO $$
BEGIN
    -- Check and create cube extension if it does not exist
    IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'cube') THEN
        CREATE EXTENSION cube;
    END IF;
    
    -- Check and create earthdistance extension if it does not exist
    IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'earthdistance') THEN
        CREATE EXTENSION earthdistance;
    END IF;
END $$;

CREATE TABLE merchants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    merchant_category VARCHAR(25) NOT NULL CHECK (merchant_category IN ('SmallRestaurant', 'MediumRestaurant', 'LargeRestaurant', 'MerchandiseRestaurant', 'BoothKiosk', 'ConvenienceStore')),
    image_url TEXT NOT NULL,
    location_lat FLOAT NOT NULL,
    location_long FLOAT NOT NULL,
    earth_location CUBE,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE INDEX idx_merchants_earth_location ON merchants USING GiST (earth_location);
