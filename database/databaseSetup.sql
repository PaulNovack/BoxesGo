CREATE TABLE users (
                       id INT AUTO_INCREMENT,
                       username TEXT NOT NULL,
                       password TEXT NOT NULL,
                       created_at DATETIME not null DEFAULT current_timestamp,
                       PRIMARY KEY (id)
);