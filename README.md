## Introducción

Desarrollando una API consultiva en Go con Gin Web Framework


### Capa de Persistencia

Para este ejemplo utilizaremos una base de datos SQLite en memoria, que se encontrará representada en la variable DB dentro del package repository.
* `repository.CreateConnection()` - Crea una conexión con una base de datos SQLite.
* `repository.CreateUser(newUser *model.User)` - Crea un nuevo usuario en la base de datos.
* `repository.GetAllUsers()` - Obtiene todos los usuarios de la base de datos.
* `repository.GetUserByID(id int)` - Obtiene un usuario con un ID específico.


### Capa Web
Usaremos Gin Framework como nuestra herramienta para manejar las peticiones recibidas por HTTP. Para este ejemplo se usaremos los métodos GET y POST
* `/api/users` - Obtiene un listado de usuarios `model.User` 
* `/api/users/[USER_ID]` -  Obtiene un usuario específico por el ID.
* `/api/users` - POST. Solicitamos la creación de un nuevo usuario enviandole como body un objeto `model.User` en formato JSON.

