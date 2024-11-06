CREATE TABLE shopping_cart
(
    id VARCHAR(255) NOT NULL,
    user_phone CHAR(13) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    quantity INT(4) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_shopping_card_user(user_phone) REFERENCES users(no_telepon),
    FOREIGN KEY fk_shopping_card_product(product_id) REFERENCES products(id)
);