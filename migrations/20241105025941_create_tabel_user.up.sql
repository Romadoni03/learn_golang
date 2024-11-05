CREATE TABLE users
(
    no_telepon CHAR(13) NOT NULL,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(200) NOT NULL,
    last_updated_username DATETIME,
    name VARCHAR(200),
    email VARCHAR(200),
    photo_profile VARCHAR(200),
    bio TEXT,
    gender VARCHAR(20),
    status_member VARCHAR(200),
    birth_date VARCHAR(200),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    token VARCHAR(255),
    token_expired_at DATETIME,
    PRIMARY KEY(no_telepon)
)ENGINE = InnoDB;