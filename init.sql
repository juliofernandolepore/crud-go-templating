CREATE DATABASE basededatos1;

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

INSERT INTO usuario(nombre, correo, password, activo) VALUES("javier lopez", "javiervito@gmail.com", "passworddebil", 1);