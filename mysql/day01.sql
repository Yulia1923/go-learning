CREATE DATABASE test_db;

USE test_db;

CREATE TABLE users (

    id INT PRIMARY KEY,

    name VARCHAR(50),

    age INT

);

CREATE TABLE orders (

    id INT PRIMARY KEY,

    user_id INT,

    product VARCHAR(50)

);

INSERT INTO users (id, name, age)

VALUE (1,'Kim',20);

INSERT INTO users (id, name, age)

VALUE (2, 'Park', 17);

INSERT INTO users (id, name, age) 

VALUE (3, 'Siwoo', 25);

INSERT INTO users (id, name, age)

VALUE (4, 'Yoonsik', 28);

INSERT INTO orders (id, user_id, product)

VALUE (101,1,'Laptop');

INSERT INTO orders (id, user_id, product)

VALUE (102,3,'Mouse');

INSERT INTO orders(id, user_id, product)

VALUE (103,1,'Keyboard');

INSERT INTO orders (id, user_id, product)

VALUE (104, 4,'Phone');

UPDATE users

SET name='Siwoo'

WHERE id = 3;

SELECT * 

FROM users 

WHERE age > 18

AND name = 'Kim'

OR name = 'Siwoo';

DELETE FROM orders;


SELECT *

FROM users

ORDER BY age ASC

LIMIT 3;

SELECT *

FROM users

ORDER BY age DESC

LIMIT 2;

SELECT *

FROM users

WHERE age IN (17,20,25,28);

SELECT *

FROM users

WHERE age BETWEEN 18 AND 30;

SELECT *

FROM users

WHERE name LIKE "S%";

SELECT *

FROM users

WHERE age > 18 

ORDER BY AGE DESC

LIMIT 3;

SELECT orders.id,users.name,orders.product

FROM orders

INNER JOIN users

ON orders.user_id = users.id

WHERE users.name = 'Siwoo';

SELECT users.name, orders.product

FROM users

LEFT JOIN orders

ON users.id = orders.user_id

WHERE orders.product = 'Laptop';