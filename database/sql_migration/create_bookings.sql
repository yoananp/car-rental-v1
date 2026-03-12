CREATE TABLE (
    id INTEGER PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES customers(id),
    car_id INT NOT NULL REFERENCES cars(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_price INT NOT NULL,
    status VARCHAR(20) CHECK (status IN ('Pending','Confirmed','Completed','Cancelled'))
);