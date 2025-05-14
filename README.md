## Requisitos

- **Go**: Versión 1.24.2 o superior.
- **Dependencias**: Las librerías necesarias están especificadas en el archivo `go.mod`.

## Librerías utilizadas

A continuación, se describen las librerías directas utilizadas en el proyecto y su utilidad:

- **`github.com/asaskevich/govalidator`**: Proporciona funciones para la validación de datos.
- **`github.com/bootcamp-go/web`**: Librería personalizada para manejar solicitudes y respuestas HTTP.
- **`github.com/go-chi/chi/v5`**: Framework ligero para construir servidores HTTP en Go.
- **`github.com/joho/godotenv`**: Permite cargar variables de entorno desde un archivo `.env`.
- **`github.com/swaggo/http-swagger`**: Proporciona una interfaz para servir la documentación Swagger en el servidor.
- **`github.com/swaggo/swag`**: Herramienta para generar documentación Swagger a partir de comentarios en el código.
- **`github.com/go-sql-driver/mysql`**: Driver para conectar aplicaciones Go con bases de datos MySQL.

## Configuración

1. **Archivo `.env`**: Crea un archivo `.env` en la raíz del proyecto con las variables necesarias para la configuración. Puedes basarte en el archivo `.env.example` proporcionado.

2. **Cargar configuración**: El archivo `cmd/config` se encarga de cargar las variables de entorno.

## Configuración de la base de datos

Para ejecutar la base de datos, asegúrate de tener Docker instalado y ejecuta el siguiente comando en la raíz del proyecto:

```bash
  docker-compose up
```
Esto levantará la base de datos definida en el archivo docker-compose.yml.

## Uso de Swagger

Swagger está integrado en el proyecto para documentar los endpoints. Para acceder a la documentación:

1. Ejecuta el servidor (ver sección "Ejecución").
2. Abre un navegador y accede a la URL: `http://localhost:8080/swagger/index.html`.
3. Verás la documentación Swagger con todos los endpoints y sus respectivas descripciones.
4. Puedes probar los endpoints directamente desde la interfaz de Swagger haciendo clic en el botón "Try it out" y luego en "Execute".
5. También puedes ver los modelos de datos y los códigos de respuesta de cada endpoint en la sección "Models" y "Responses", respectivamente.

### Creación de documentación Swagger

La documentación Swagger se genera automáticamente a partir de los comentarios en el código. Para agregar documentación a un endpoint, sigue estos pasos:

1. Abre el archivo correspondiente al endpoint (por ejemplo, `internal/app/warehouse/handler.go`).
2. Agregar la documentación correspondiente al método del endpoint. Esto se deberá hacer teniendo en cuenta la documentación que swaggo https://github.com/swaggo/swag
3. Asegúrate de que el método del endpoint tenga los tags `@Summary` y `@Description` para que se muestren correctamente en Swagger.
4. Ejecutar el comando `swag init -g cmd/main.go` para generar la documentación Swagger actualizada. Esto generará un archivo `docs/swagger.json` que será utilizado por el servidor para mostrar la documentación.

## Ejecución

Sigue estos pasos para ejecutar el proyecto:

1. **Instalar dependencias**:
   ```bash
   go mod tidy