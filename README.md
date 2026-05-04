# agropecuario_crud
 
El API provee la gestión de los diferentes procesos relacionados con el sector agropecuario, permitiendo la administración de categorías de ganado, subastas y precios de subasta del sistema.
 
## Especificaciones Técnicas
 
### Tecnologías Implementadas y Versiones
 
* __[Golang](https://go.dev/doc/install)__
* __[Gorilla Mux](https://github.com/gorilla/mux)__
### Variables de Entorno
 
```
AGROPECUARIO_CRUD_PGDB=[nombre de la base de datos]
AGROPECUARIO_CRUD_PGPASS=[password del usuario]
AGROPECUARIO_CRUD_PGURLS=[direccion de la base de datos]
AGROPECUARIO_CRUD_PGPORT=[Puerto de conexión con la base de datos]
AGROPECUARIO_CRUD_PGUSER=[usuario con acceso a la base de datos]
AGROPECUARIO_CRUD_PGSCHEMA=[esquema donde se ubican las tablas]
AGROPECUARIO_CRUD_HTTP_PORT=[puerto de ejecucion]
```
 
NOTA: Las variables se pueden ver en el fichero `config/db.go` y están identificadas con `AGROPECUARIO_CRUD_...`
 
### Ejecución del Proyecto
 
```
#1. Obtener el repositorio con Go
go get github.com/tu_usuario/agropecuario_crud
 
#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/tu_usuario/agropecuario_crud
 
#3. Moverse a la rama develop
git pull origin develop && git checkout develop
 
#4. Alimentar todas las variables de entorno que utiliza el proyecto.
AGROPECUARIO_CRUD_HTTP_PORT=8082 AGROPECUARIO_CRUD_PGHOST=127.0.0.1 AGROPECUARIO_CRUD_PGPORT=5432 AGROPECUARIO_CRUD_PGUSER=postgres AGROPECUARIO_CRUD_PGPASS=postgres AGROPECUARIO_CRUD_PGDB=AgroCampo AGROPECUARIO_CRUD_PGSCHEMA=agropecuario go run main.go
```
 


