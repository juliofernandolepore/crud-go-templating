# crud-go-templating

dabatabase conection

mysql (installing drivers)
```
$ go get -u github.com/go-sql-driver/mysql
```
third lib 

go get github.com/joho/godotenv

i use .env but not include in gitignore for learning purpose

# docker-steps

```
docker run -d -p 3306:3306 --name contenedor-mysql -e MYSQL_ROOT_PASSWORD=PassWord123. mysql
```
get conection from terminal inside container

```
sudo docker exec -it contenedor-mysql mysql -p
```
password: PassWord123.

# Go run

go run main.go

``````
http://localhost:3000/users/
```