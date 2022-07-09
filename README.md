# Golang api user

El siguiente ejersicio cumple con las tareas:

1. Crear un servicio de registro de usuario que reciba como parámetros usuario, correo,
teléfono y contraseña.
2. El servicio deberá validar que el correo y telefono no se encuentren registrados, de lo
contrario deberá retornar un mensaje “el correo/telefono ya se encuentra registrado”.
3. Deberá validar que la contraseña sea de 6 caracteres mínimo y 12 máximo y contener
al menos una mayúscula, una minúscula, un carácter especial (@ $ o &amp;) y un número.
4. Validar que el teléfono sea a 10 dígitos y el correo tenga un formato válido.
5. Crear un servicio login que reciba como parámetros usuario o correo y contraseña.
6. El servicio debe devolver un token jwt.
7. Deberá validar que el usuario o correo y contraseña sean válidos, de lo contrario
retorna un mensaje “usuario / contraseña incorrectos”.
8. En ambos servicios se deberá validar que todos los parámetros solicitados vayan en el
cuerpo de la petición, de lo contrario retorna un mensaje con el campo faltante.

## Dependencias

Para iniciar el proyecto se requiere de Postgresql

## Iniciar el proyecto

Para iniciar el proyecto se puede realizar por medio de docker, utilizando composer, o realizar a mano, solo teniendo encuenta las dependencias como lo es postgres.

## Configuracion

*Nota*: las variables son requeridas si el proyecto es iniciado a mano.

#### Api

- DB_DATABASE=db // Nombre de la base de datos
- DB_HOST=postgres // Host del servidor
- DB_PASSWORD=password // Password
- DB_USER=usuario // Usuario de la db

## Docker-composer

Para iniciar utilizando docker, basta con ejecutar el siguiente comando en la raiz del proyecto, sin requerir entrar a una carpeta especifica (api, reto).

```
docker-compose up --remove-orphans
```

se recomienda iniciarlo con docker, para evitar instalar de forma manual las dependencias del codigo o las DB.

## Direcciones

El proyecto expone las siguientes url

- POST http://localhost:3000/user
- POST http://localhost:3000/login
- GET http://localhost:3000/me

### Registro de usuario

Para el registro del usuario se utiliza la direccion http://localhost:3000/user pasando por POST los siguientes datos.

```
{
    "username": "usuario",
    "phone": "telefono",
    "email": "correo",
    "password": "eXixo2@ss"
}
```

### Login

El Login utiliza la direccion POST http://localhost:3000/login por POST, se puede pasar el usuario o correo, ejemmplo:

```
{
    "username": "usuario",
    "password": "eXixo2@ss"
}
```

esto nos regresa un token de sesion, que podremos utilizar para recuperar los datos del usuario.

### Datos del usuario

Siempre y cuando se tenga el token de la sesion, se puede utilizar la siguiente direccion http://localhost:3000/me por GET para recuperar los datos,
para ello se debe pasar el token por la cabecera Authorization, no requiere nada mas, que el token puro.
