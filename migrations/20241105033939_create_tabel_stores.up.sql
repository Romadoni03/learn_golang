CREATE TABLE stores
(
    store_id VARCHAR(255) NOT NULL,
    no_telepon CHAR(13) NOT NULL,
    name VARCHAR(50) NOT NULL,
    last_updated_name DATETIME NOT NULL,
    logo VARCHAR(50) NOT NULL,
    description TEXT,
    status VARCHAR(200),
    link_store VARCHAR(100),
    total_comment INT(4),
    total_following INT(4),
    total_follower INT(4),
    total_product INT (4),
    conditions VARCHAR(200),
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY(store_id),
    Foreign Key fk_users_stores(no_telepon) REFERENCES users (no_telepon)
)ENGINE=InnoDB;