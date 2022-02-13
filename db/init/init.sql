CREATE DATABASE IF NOT EXISTS todoapp;

CREATE TABLE users (
    id int NOT NULL AUTO_CREMENT PRIMARY,
    name varchar(32) NOT NULL
);

INSERT INTO users (id,name) VALUES ("","naoki_hirahara");

SHOW DATABASES;
SHOW TABLES FROM todoapp;