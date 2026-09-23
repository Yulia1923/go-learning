CREATE DATABASE IF NOT EXISTS ecommerce_demo;

USE ecommerce_demo;

DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE products (
    id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL
);

CREATE TABLE orders (
    id INT PRIMARY KEY,
    user_id INT NOT NULL,
    status VARCHAR(20) NOT NULL,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
);

CREATE TABLE order_items (
    id INT PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,

    FOREIGN KEY (order_id)
        REFERENCES orders(id),
    
    FOREIGN KEY (product_id)
        REFERENCES products(id)
);

INSERT INTO users (id, username, email)
VALUES
(101, 'Alice', 'alice@test.com'),
(102, 'Bob', 'bob@test.com'),
(103, 'Charlie', 'charlie@test.com');

INSERT INTO products(id, name, price, stock)
VALUES
(1, 'iPhone', 5000.00, 20),
(2, 'iPad', 3000.00, 30),
(3, 'AirPods', 1000.00, 50),
(4, 'Keyboard', 500.00, 100),
(5, 'Monitor', 2000.00, 40);

INSERT INTO orders (id, user_id, status)
VALUES
(1001, 101, 'PAID'),
(1002, 102, 'PAID'),
(1003, 101, 'PENDING'),
(1004, 103, 'PAID');

INSERT INTO order_items (id, order_id, product_id, quantity)
VALUES
(1, 1001, 1, 1),
(2, 1001, 3, 2),
(3, 1002, 2, 1),
(4, 1002, 4, 2),
(5, 1003, 5, 1),
(6, 1004, 1, 1),
(7, 1004, 4, 1);

SELECT *
FROM users;

SELECT *
FROM products;

SELECT *
FROM orders;

SELECT *
FROM order_items;
