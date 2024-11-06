CREATE TABLE order_item
(
    id VARCHAR(255) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    quantity INT(4) NOT NULL,
    price DECIMAL(15, 2) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_orders_order_item(order_id) REFERENCES orders(id),
    FOREIGN KEY fk_product_item(product_id) REFERENCES products(id)
);