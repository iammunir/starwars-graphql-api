CREATE TABLE IF NOT EXISTS starships (
    id SERIAL,
    name VARCHAR(255),
    model VARCHAR(100),
    manufacturer VARCHAR(255),
    cost_in_credits BIGINT,
    length FLOAT,
    max_atmosphering_speed INT,
    crew INT,
    passengers INT,
    cargo_capacity BIGINT,
    consumables VARCHAR(20),
  	hyperdrive_rating FLOAT,
  	mglt INT,
    starship_class VARCHAR(50),
    PRIMARY KEY (id)
);