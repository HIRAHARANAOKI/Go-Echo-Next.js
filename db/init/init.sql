CREATE DATABASE IF NOT EXISTS todoapp;

CREATE TABLE users (
    id int NOT NULL PRIMARY,
    name varchar(32) NOT NULL
);

INSERT INTO users (id,name) VALUES ("","naoki_hirahara");
INSERT INTO users (id,name) VALUES ("","futa_matuo");
INSERT INTO users (id,name) VALUES ("","docker");

SHOW DATABASES;
SHOW TABLES FROM todoapp;