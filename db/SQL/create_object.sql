CREATE TABLE sample_db.category(
    id INT NOT NULL AUTO_INCREMENT,
    obj_id VARCHAR(36) NOT NULL,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_obj_id (obj_id)
);

CREATE TABLE sample_db.product(
    id INT NOT NULL AUTO_INCREMENT,
    obj_id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL, 
    price INT NOT NULL,
    category_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_obj_id (obj_id),
    FOREIGN KEY category_fk (category_id) REFERENCES category (obj_id)
);