ALTER TABLE transactions
    ADD COLUMN store_id INT,
    ADD CONSTRAINT fk_store_id FOREIGN KEY (store_id) REFERENCES stores(store_id);