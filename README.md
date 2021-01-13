# Golang-REST-API
A REST API made in Golang using Fasthttp and Fasthttp Router

## how to run
To run you need to:
* change user to superuser: `sudo bash`
* create this folder: `mkdir -p /root/docker/mariadb/var/lib/mysql`
* clone this repository: `git clone https://github.com/TiagoCavalcanteTrindade/Golang-REST-API`
* go to directory: `cd Golang-REST-API`
* create a file called `.env` and put inside it the following `env`s:
	* **MYSQL_ROOT_PASSWORD**: MariaDB password
	* **PASSWORD**: the same than `MYSQL_ROOT_PASSWORD`
	* **USER**: MariaDB user
	* **DATABASE**: MariaDB database name
* lift the MariaDB container: `docker-compose up mariadb -d`
* install MariaDB client: `sudo apt install mysql-client -y`
* connect to MariaDB: `mysql -uMARIADB_USER_HERE -pMARIADB_PASSWORD_HERE -hYOUR_IP_HERE`
* create a database: `CREATE DATABASE your_database_name_here;`
* use this database: `USE your_database_name_here`
* create this table: `CREATE TABLE users (id INT PRIMARY KEY AUTO_INCREMENT, name VARCAHR(64), email VARCHAR(128));`
* exit of MariaDB client: `exit`
* set on your Shell the `env` **PORT**: `export PORT=your_port`
* lift the REST-API container: `docker-compose up rest-api -d`

## routes
| method | path          | description |
|--------|---------------|-------------|
| GET    | `/users`      | return all users and their information
| GET    | `/users/<id>` | return the information about the user with the id `id`
| POST   | `/users`      | create a user, receives a `application/json` body with the required parameters `name` and `email`, return the created user's id
| PUT    | `/users/<id>` | edit the information about the user with the id `id`, receives a `application/json` body with the required parameters `name` and `email`
| DELETE | `/users/<id>` | delete the user with the id `id`