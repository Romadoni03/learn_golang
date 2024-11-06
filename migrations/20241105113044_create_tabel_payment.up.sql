CREATE TABLE payment
(
    id VARCHAR(255) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    payment_method ENUM('e-wallet','virtual-account','transfer-bank') NOT NULL,
    status ENUM('success','pending','failed') NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_payment_order(order_id) REFERENCES orders(id)
);