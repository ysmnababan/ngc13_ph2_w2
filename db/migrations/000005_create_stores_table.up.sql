CREATE TABLE IF NOT EXISTS stores (
    store_id serial primary key,
    store_name varchar(255),
    store_address varchar(255),
    longitude decimal(10,2),
    latitude decimal(10,2),
    rating int
);