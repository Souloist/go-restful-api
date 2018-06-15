CREATE TABLE addresses (
    id          SERIAL PRIMARY KEY, 
    street      TEXT NOT NULL, 
    city        TEXT NOT NULL, 
    state       TEXT NOT NULL, 
    country     TEXT NOT NULL, 
    zip_code    TEXT NOT NULL
); 

CREATE TABLE persons ( 
    id         SERIAL PRIMARY KEY, 
    first_name TEXT NOT NULL, 
    last_name  TEXT NOT NULL
); 

CREATE TABLE person_addresses ( 
    person_id  INTEGER REFERENCES persons, 
    address_id INTEGER REFERENCES addresses, 
    PRIMARY KEY (person_id, address_id)
);
