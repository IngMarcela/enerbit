# CRUD Enerbit en Golang
Este proyecto es una API-Rest implementada en Golang que se enfrenta a un desafio de registrar y realizar un seguimiento
eficiente de las ordenes de servicio para un predio o cliente en particular.
## Tecnologías utilizadas

- Golang
- MySQL (motor de base de datos)
- GORM (ORM para interactuar con la base de datos)
- OpenAPI 3 Swagger (documentación autogenerada)


## Funcionalidades
La API cuenta con los siguientes endpoints:

- GET /customers/active: Devuelve una lista de clientes que se encuentran activos.
- POST /customers/store: Crea un nuevo cliente en la base de datos con los campos first name, lastname, address.
- POST /workOrder/store: Crea un nueva orden ligado a un cliente en la base de datos.
- GET /workOrder/date: Devuelve una lista de ordenes activas.
- GET /workOrder/customer/{id}: Devuelve la información del cliente con el id de la orden
- GET /workOrder/customer/{id}: Devuelve la información del cliente y la orden con el id de la orden

## Documentación
La documentación de la API se encuentra disponible en formato OpenAPI 3 Swagger y puede ser consultada mediante el endpoint GET /swagger.

## Despliegue

1. Clonar este repositorio.
2. correr la aplicacion. 