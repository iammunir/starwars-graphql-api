CREATE TABLE IF NOT EXISTS species (
  	id SERIAL,
    name VARCHAR(100),
    classification VARCHAR(100),
    designation VARCHAR(100),
    average_height SMALLINT,
    skin_colors VARCHAR(100),
    hair_colors VARCHAR(100),
    eye_colors VARCHAR(100),
    average_lifespan SMALLINT,
    language VARCHAR(100),
    homeworld_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (homeworld_id) REFERENCES planets(id)
);