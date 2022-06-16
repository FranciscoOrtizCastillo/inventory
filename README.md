
# Ejemplo de CRUD en GO. Control de Inventario.

Organización del proyecto usando Clean Architecture.

repository -> service -> api

# Pasos de la creación del proyecto

```
git init .

go mod init github.com/FranciscoOrtizCastillo/inventory

go get gopkg.in/yaml.v3
go get go.uber.org/fx
go get github.com/jmoiron/sqlx
go get github.com/go-sql-driver/mysql

```

# Uso de MariaDB con Docker

````
docker pull mariadb:10.7.4

docker image ls

# Crear un contenedor con mariaDB
docker run -d --name mariadb -p 3306:3306 --env MYSQL_ROOT_PASSWORD=secret#password123 --env MYSQL_DATABASE=inventory mariadb:10.7.4 

docker logs mariadb

docker ps

docker stop mariadb

docker start mariadb
```

