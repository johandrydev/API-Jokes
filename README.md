API de Chistes
==============

Esta API proporciona una manera divertida de obtener chistes aleatorios.

Instrucciones
-------------

Para utilizar este Repositorio, solo debes seguir los siguientes pasos:

1.  Clona este repositorio en tu computadora.
2.  Instala las dependencias con el comando `go mod tidy`.
3.  Ejecuta el servidor con el comando `go run main.go`.
4.  Abre tu navegador y navega a la ruta `http://localhost:8000/api/jokes`.
5.  Tambien puedes ejecutar el comando `go build .` para compilar el proyecto y luego ejecutar el archivo binario generado.

Para utilizar esta API, solo debes hacer una petición GET a la ruta `/api/jokes`. La respuesta será 25 chistes aleatorio en formato JSON.:

Rutas
-----

### `GET /api/jokes`

Devuelve 25 chistes aleatorio en formato JSON.

#### Parámetros de consulta

No hay parámetros de consulta necesarios para esta ruta.

#### Respuesta

-   `200 OK` en caso de éxito. El cuerpo de la respuesta contendrá un objeto JSON con la siguiente estructura:

jsonCopy code

```
{
	"status": "success",
	"data": [
		{
			"id": "Afs59AUXTzWQepoCo63QZw",
			"url": "https://api.chucknorris.io/jokes/Afs59AUXTzWQepoCo63QZw",
			"value": "The U.S. Military once tried to capture the power of a round house kick from Chuck Norris into a bomb. It was called the Manhattan Project and it didn't even come close.",
			"icon_url": "https://assets.chucknorris.host/img/avatar/chuck-norris.png"
		},
		{
			"id": "YblD05tGRV2Bw2oAqItvDQ",
			"url": "https://api.chucknorris.io/jokes/YblD05tGRV2Bw2oAqItvDQ",
			"value": "When Chuck Norris plays Five Nights at Freddy's, the animatronics don't scare him. The animatronics come to him",
			"icon_url": "https://assets.chucknorris.host/img/avatar/chuck-norris.png"
		},
  ]
}
```

-   `400 Bad Request` en caso de que ocurra un error en la peticion.
-   `500 Internal Server Error` en caso de que ocurra un error en el servidor.