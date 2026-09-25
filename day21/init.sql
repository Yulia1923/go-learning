CREATE DATABASE IF NOT EXISTS go_mysql_demo;

USE go_mysql_demo;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);

INSERT INTO users (name, age)
VALUES
    ('Alice', 20),
    ('Bob', 21),
    ('Charlie', 22);

    