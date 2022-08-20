
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

#Instalar todas las dependencias
go get ./...   

# o
go mod tidy

```

# Uso de MariaDB con Docker

```
docker pull mariadb:10.7.4

docker image ls

# Crear un contenedor con mariaDB
docker run -d --name mariadb -p 3306:3306 --env MYSQL_ROOT_PASSWORD=secret#password123 --env MYSQL_DATABASE=inventory mariadb:10.7.4 

docker logs mariadb

docker ps

docker stop mariadb

docker start mariadb
```

## mockery

https://github.com/vektra/mockery

````
go install github.com/vektra/mockery/v2@latest

brew install mockery
brew upgrade mockery

go generate ./...

go get github.com/stretchr/testify/mock


```

https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/09.6.html



## Añadir roles

```
go generate ./...

```

## Añadir productos

1.- Crear metodos en repository 

    SaveProduct
	GetProducts
	GetProductByID

2.- Implementar los metodos en products.repository y crear entity product

3.- Crear model products

4.- Implementar los metodos en products.service 

```
go generate ./...

```

5.- Crear unit test para el servicio