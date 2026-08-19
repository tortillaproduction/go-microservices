CREATE TABLE sample_db.category(
    internal_id INT NOT NULL AUTO_INCREMENT,
    id VARCHAR(36) NOT NULL,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY (internal_id),
    UNIQUE KEY idx_id (id)
);

CREATE TABLE sample_db.product(
    internal_id INT NOT NULL AUTO_INCREMENT,
    id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL, 
    price INT NOT NULL,
    category_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (internal_id),
    UNIQUE KEY idx_id (id),
    FOREIGN KEY category_fk (category_id) REFERENCES category (id)
);