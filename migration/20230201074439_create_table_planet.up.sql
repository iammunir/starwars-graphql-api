CREATE TABLE IF NOT EXISTS planets (
    id SERIAL,
    name VARCHAR(100),
    rotation_period SMALLINT,
    orbital_period SMALLINT,
    diameter INT,
    climate VARCHAR(100),
    gravity VARCHAR(100),
    terrain VARCHAR(100),
    surface_water SMALLINT,
    population BIGINT,
    PRIMARY KEY (id)
);