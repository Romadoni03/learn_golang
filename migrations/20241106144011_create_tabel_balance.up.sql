CREATE TABLE balance
(
    id VARCHAR(255) NOT NULL,
    user_phone CHAR(13) NOT NULL,
    current_balance DECIMAL(15,2) DEFAULT 0,
    last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY fk_balance_user(user_phone) REFERENCES users(no_telepon)
);