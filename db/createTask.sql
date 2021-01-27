/*
専用のユーザ(todo)を作成し権限を付与しておく
$ mysql -u root -p
MariaDB [(none)]> create database tododb default charset utf8;
MariaDB [(none)]> create user 'todo'@'localhost' identified by 'todo';
MariaDB [(none)]> grant all on 'tododb'.* to 'todo'@'localhost';
*/
USE tododb
DROP TABLE IF EXISTS task;
CREATE TABLE IF NOT EXISTS task {
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
};