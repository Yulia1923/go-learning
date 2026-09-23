CREATE DATABASE IF NOT EXISTS shop_demo;

USE shop_demo;

DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
    id INT PRIMARY KEY,
    user_id INT,
    product VARCHAR(50),
    amount DECIMAL(10,2)
);

INSERT INTO orders (id, user_id, product, amount)
VALUES
(1, 101, 'iphone', 5000),
(2, 102, 'ipad',3000),
(3, 101, 'Airpods', 1000),
(4, 103, 'Macbook', 8000),
(5, 102, 'Mouse', 200),
(6, 101, 'Keyboard', 500),
(7, 103, 'Monitor', 2000),
(8, 102, 'Headset', 800);

SELECT *
FROM orders;

SELECT COUNT(*) AS total_orders
FROM orders;

SELECT SUM(amount) AS total_sales
FROM orders;

SELECT AVG(amount) AS average_order
FROM orders;

SELECT MAX(amount) AS max_order
FROM orders;

SELECT MIN(amount) AS min_order
FROM orders;

SELECT 
    user_id,
    COUNT(*) AS order_count
FROM orders
GROUP BY user_id;

SELECT 
    user_id,
    SUM(amount) AS total_spent
FROM orders
GROUP BY user_id;

SELECT 
    user_id,
    AVG(amount) AS average_order
FROM orders
GROUP BY user_id;

SELECT 
    user_id,
    MAX(amount) AS max_order
FROM orders
GROUP BY user_id;

SELECT 
    user_id,
    MIN(amount) AS min_order
FROM orders
GROUP BY user_id;

SELECT 
    user_id,
    COUNT(*) AS order_count
FROM orders
GROUP BY user_id
HAVING COUNT(*) >2;

SELECT 
    user_id,
    SUM(amount) AS total_spent
FROM orders
GROUP BY user_id
HAVING SUM(amount) > 5000;

SELECT 
    user_id,
    COUNT(*) AS order_count,
    SUM(amount) AS total_spent
FROM orders
WHERE amount >= 1000
GROUP BY user_id;

SELECT
    user_id,
    COUNT(*) AS order_count,
    SUM(amount) AS total_Spent
FROM orders
WHERE amount >= 1000
GROUP BY user_id
HAVING COUNT(*) >= 2;

SELECT
    user_id,
    COUNT(*) AS order_count,
    SUM(amount) AS total_spent,
    AVG(amount) AS average_order,
    MAX(amount) AS max_order,
    MIN(amount) AS min_order
FROM orders
GROUP BY user_id;

SELECT
    user_id,
    COUNT(*) AS order_count,
    SUM(amount) AS total_spent,
    AVG(amount) AS average_order
FROM orders
GROUP BY user_id
HAVING SUM(amount) > 5000;


