CREATE TABLE shipping
(
    id VARCHAR(255) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    shipping_address TEXT NOT NULL,
    city VARCHAR(100) NOT NULL,
    postal_code INT(8) NOT NULL,
    province VARCHAR(100) NOT NULL,
    phone_number CHAR(13) NOT NULL,
    shipping_cost DECIMAL(10,2) NOT NULL,
    shipping_status ENUM('pending','shipped','delivered'),
    tracking_number VARCHAR(100),
    shipped_date DATETIME,
    delivered_date DATETIME,
    PRIMARY KEY(id),
    FOREIGN KEY fk_shipping_order(order_id) REFERENCES orders(id)
);