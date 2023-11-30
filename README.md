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