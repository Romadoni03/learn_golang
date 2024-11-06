CREATE TABLE wishlists
(
    id VARCHAR(255) NOT NULL,
    user_phone CHAR(13) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_wishlist_user(user_phone) REFERENCES users(no_telepon),
    FOREIGN KEY fk_wishlist_product(product_id) REFERENCES products(id)
);