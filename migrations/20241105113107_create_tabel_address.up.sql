CREATE TABLE addresses
(
    id VARCHAR(255) NOT NULL,
    user_phone CHAR(13) NOT NULL,
    address_line VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    zip_code INT(8),
    created_at DATETIME NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_address_user(user_phone) REFERENCES users(no_telepon)
);