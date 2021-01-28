/*
専用のユーザ(todo)を作成し権限を付与しておく
$ mysql -u root -p
mysql> create database tododb default charset utf8;
mysql> create user 'todo'@'localhost' identified by 'todo';
mysql> grant all on `tododb`.* to 'todo'@'localhost';
*/
USE tododb
DROP TABLE IF EXISTS task;
CREATE TABLE IF NOT EXISTS task (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY(id)
);