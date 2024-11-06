CREATE TABLE transaction
(
    id VARCHAR(255) NOT NULL,
    user_phone CHAR(13) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    payment_id VARCHAR(255) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    status ENUM('pending','completed','failed'),
    PRIMARY KEY(id),
    FOREIGN KEY fk_transaction_users(user_phone) REFERENCES users(no_telepon),
    FOREIGN KEY fk_transaction_order(order_id) REFERENCES orders(id),
    FOREIGN KEY fk_transaction_payment(payment_id) REFERENCES payment(id)
);