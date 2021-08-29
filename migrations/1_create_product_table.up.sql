CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(40),
    description TEXT,
    price FLOAT,
) engine = InnoDB

DEFAULT charset = utf8;