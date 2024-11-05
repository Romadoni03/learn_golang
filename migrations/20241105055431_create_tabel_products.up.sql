CREATE TABLE products
(
    id VARCHAR(255) NOT NULL,
    store_id VARCHAR(255) NOT NULL,
    photo_product VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    dangerius_product VARCHAR(50) NOT NULL,
    price DECIMAL(15, 2) NOT NULL,
    stock INT(8) NOT NULL,
    wholesaler VARCHAR(200) NOT NULL,
    shipping_cost DECIMAL(15, 2) NOT NULL,
    shipping_insurance VARCHAR(200) NOT NULL,
    conditions VARCHAR(255) NOT NULL,
    pre_order VARCHAR(200) NOT NULL,
    status VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_stores_products(store_id) REFERENCES Stores(store_id)
)ENGINE=InnoDB;