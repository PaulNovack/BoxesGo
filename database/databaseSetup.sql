CREATE TABLE users (
                       id INT AUTO_INCREMENT,
                       username TEXT NOT NULL,
                       password TEXT NOT NULL,
                       created_at DATETIME not null DEFAULT current_timestamp,
                       PRIMARY KEY (id)
);

ALTER TABLE `boxes`.`users`
    ADD COLUMN `authkey` VARCHAR(45) NULL AFTER `password`;
