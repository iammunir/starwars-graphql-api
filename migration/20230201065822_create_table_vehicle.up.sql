CREATE TABLE IF NOT EXISTS vehicles (
    id SERIAL,
    name VARCHAR(255),
    model VARCHAR(100),
    manufacturer VARCHAR(255),
    cost_in_credits INT,
    length FLOAT,
    max_atmosphering_speed INT,
    crew INT,
    passengers INT,
    cargo_capacity INT,
    consumables VARCHAR(20),
    vehicle_class VARCHAR(50),
    PRIMARY KEY (id)
);