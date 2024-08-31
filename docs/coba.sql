create table users (
    user_id VARCHAR (255) PRIMARY KEY,
    no_telepon CHAR (13) UNIQUE,
    password VARCHAR (255),
    username VARCHAR (200),
    last_updated_username DATETIME,
    name VARCHAR (200) UNIQUE,
    email VARCHAR (200),
    photo_profile VARCHAR (200),
    bio TEXT,
    gender ENUM('laki-laki','perempuan'),
    status_member VARCHAR (200),
    birth_date DATETIME,
    created_at TIMESTAMP,
    token VARCHAR (255),
    token_expired_at BIGINT
);

CREATE Table Stores (
    store_id VARCHAR (255) PRIMARY KEY NOT NULL,
    user_id VARCHAR (255) NOT NULL,
    name VARCHAR (255) UNIQUE NOT NULL,
    last_updated_name DATETIME NOT NULL,
    logo VARCHAR (200),
    description TEXT,
    status VARCHAR(200) NOT NULL,
    link_store VARCHAR (100) NOT NULL,
    total_comment INT (4) NOT NULL,
    total_following INT (4) NOT NULL,
    total_follower INT (4) NOT NULL,
    total_product INT (4) NOT NULL,
    conditions VARCHAR (200) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    Foreign Key fk_users_stores(user_id) REFERENCES users (user_id)
)ENGINE=InnoDB;

CREATE TABLE products (
    product_id VARCHAR (255) PRIMARY KEY NOT NULL,
    store_id VARCHAR (255) NOT NULL,
    photo_product VARCHAR (255) NOT NULL,
    name VARCHAR (255) NOT NULL,
    category VARCHAR (255) NOT NULL,
    description TEXT NOT NULL,
    dangerius_product VARCHAR(255) NOT NULL,
    price DECIMAL(15, 2) NOT NULL,
    stock INT (8) NOT NULL,
    wholesaler VARCHAR (200) NOT NULL,
    shipping_cost DECIMAL (15, 2) NOT NULL,
    shipping_insurance VARCHAR (200) NOT NULL,
    conditions VARCHAR(255) NOT NULL,
    pre_order VARCHAR (200) NOT NULL,
    status VARCHAR (200) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_updated_at DATETIME NOT NULL,

    FOREIGN KEY fk_stores_products(store_id) REFERENCES Stores(store_id)
);
