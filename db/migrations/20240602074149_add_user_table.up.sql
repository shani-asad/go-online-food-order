CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(35) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX sdg ON users (id);
CREATE INDEX fddsm ON users (email);
CREATE INDEX gukut ON users (username);
CREATE INDEX gfdsasgfd ON users (created_at);
