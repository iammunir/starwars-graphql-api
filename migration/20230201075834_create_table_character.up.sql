CREATE TABLE IF NOT EXISTS characters (
    id SERIAL,
    name VARCHAR(100),
    height SMALLINT,
    mass SMALLINT,
    hair_color VARCHAR(100),
    skin_color VARCHAR(100),
    eye_color VARCHAR(100),
    birth_year VARCHAR(20),
    gender VARCHAR(20),
    homeworld_id INT,
    species_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (homeworld_id) REFERENCES planets (id),
    FOREIGN KEY (species_id) REFERENCES species (id)
);