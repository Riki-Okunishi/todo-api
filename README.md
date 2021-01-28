# todo-api

TODO管理のAPI

## REST

| Path | HTTP Method |
| ---  | --- |
| /tasks | GET |
| /tasks | POST |
| /tasks/:id | PATCH |
| /tasks/:id | DELETE |

## Dependency
+ [gin](https://github.com/gin-gonic/gin)
+ [mysql](https://www.mysql.com/jp/)
+ [Nginx](https://www.nginx.com/)

## Setup
MySQLサーバの設定
```bash
$ mysql -u root -p
mysql> create database tododb default charset utf8;
mysql> create user 'todo'@'localhost' identified by 'todo';
mysql> grant all on `tododb`.* to 'todo'@'localhost';
mysql >set global sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION'
mysql> set password for 'todo'@'localhost'="XXX"
mysql> source db/createTask.sql
```

Nginxのリパースプロキシの設定(`etc/nginx/sites-available/default`)
```
location ^~ /todo/ {
  proxy_pass  http://127.0.0.1:9000;
}
```

サーバの起動
```bash
$ MYSQL_TODO_PASSWORD=XXX go run src/main.go
```
