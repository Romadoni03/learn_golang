CREATE TABLE balance_transaction
(
    id VARCHAR(255) NOT NULL,
    balance_id VARCHAR(255) NOT NULL,
    transaction_type ENUM('credit', 'debit') NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    status ENUM('pending','completed','failed'),
    PRIMARY KEY(id),
    FOREIGN KEY fk_transaction_balance(balance_id) REFERENCES balance(id)
);