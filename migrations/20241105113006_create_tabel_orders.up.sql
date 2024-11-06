CREATE TABLE orders
(
    id VARCHAR(255) NOT NULL,
    user_phone CHAR(13) NOT NULL,
    total_aomunt DECIMAL(15,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_user_orders(user_phone) REFERENCES users (no_telepon)
)ENGINE=InnoDB;