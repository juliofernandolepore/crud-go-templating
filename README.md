# crud-go-templating

dabatabase conection

mysql (installing drivers)

$ go get -u github.com/go-sql-driver/mysql

third lib 

go get github.com/joho/godotenv

i use .env but not include in gitignore for learning purpose

# docker-steps

create and run docker container with mysql image, ports and env , password too

example:

docker run -d -p 33060:3306 --name MysqlContainer -e MYSQL_ROOT_PASSWORD=PassWord123. mysql

get conection from terminal inside container

docker exec -it MysqlContainer mysql -p

password

CREATE DATABASE basededatos1;

# sql-commands(MYSQL)

use basededatos1;

CREATE TABLE usuario (
    id int(99) NOT NULL AUTO_INCREMENT, 
    nombre varchar(200), 
    correo varchar(150), 
    password varchar(255), 
    activo int(10), 
    PRIMARY KEY (id)
    );

SHOW TABLES;

SHOW COLUMNS FROM usuario;

SELECT * FROM usuario;

#example-test-table-practice-sql

INSERT INTO usuario(nombre, correo, password, activo) VALUES("javier lopez", "javiervito@gmail.com", "passworddebil", 1);

