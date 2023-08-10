# Generador de códigos TOTP en la línea de comandos

## Pasos previos

Instalar las dependencias utlizando `go get`

## Ejecutar

### Usando código fuente

```go
go run cmd/2fa.go
```

### Utilizando el binario para linux

```sh
./bin/2fa
```

### Utilizando la versión WEB

Agregar el parámetro `-w` a la línea de comandos

## Establecer Clave

En la tabla `security` de la base de datos SQLITE3, colocar la contraseña hasheada en MD5 en el campo passwd, la clave por defecto es `123456`

## Agregar elementos a la tabla TOTP

La tabla presenta 3 columnas, `id`, `service` y `secret`

- `id`: Número de referencia, autoincrementado
- `service`: Nombre del servicio (referencial)
- `secret`: Código utilizado para generar la contraseña OTP

Por Ejemplo:
| id | service | secret |
|:--:|:-------:|:------:|
| 1  | Google  |WSXD VFRT BGTR SWYR GFDE|
