CREATE TABLE (
    id INTEGER PRIMARY KEY,
    brand VARCHAR(50) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('SUV','Sedan','MPV','Hatchback','Crossover','Wagon','Pickup','Coupe','Van','Convertible')),
    transmission VARCHAR(20) NOT NULL CHECK (transmission IN ('Manual','Automatic')),
    plate_number VARCHAR(20) NOT NULL UNIQUE,
    price_per_day INT NOT NULL,
    available BOOLEAN NOT NULL
);