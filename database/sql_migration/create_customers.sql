CREATE TABLE customers(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nik VARCHAR(20) NOT NULL UNIQUE,
    phone_number VARCHAR(15) NOT NULL

);
