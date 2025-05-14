package utils

import "errors"

// General error
var (
	ErrInternalServer = errors.New("Ha ocurrido un error, contacte al administrador")
	ErrBadRequest     = errors.New("El contenido es inválido o está mal formado")
	ErrValidation     = errors.New("Ha ocurrido un error de validación")
	ErrInvalidFilter  = errors.New("Filtro inválido: se debe proporcionar al menos un campo de filtro")
	ErrMandatoryId    = errors.New("El ID es obligatorio")
	ErrInvalidId      = errors.New("El ID debe ser un número entero válido")
)

// Warehouses errors
var (
	ErrWarehouseCodeAlreadyExists = errors.New("El código del almacén ya existe")
	ErrWareHouseNotFound          = errors.New("El almacén no existe")
	ErrWarehouseAlreadyExists     = errors.New("El almacén ya existe")
)

// Buyers errors
var (
	ErrBuyerNotFound                = errors.New("El comprador no existe")
	ErrBuyerCardNumberAlreadyExists = errors.New("El número de tarjeta ya existe")
)

// Sections errors
var (
	ErrSectionNotFound      = errors.New("No se encontró una sección asociada al ID proporcionado")
	ErrSectionAlreadyExists = errors.New("El número de sección proporcionado pertenece a una sección existente")
)

// Sellers errors
var (
	ErrSellerNotFound = errors.New("El vendedor no existe")
	ErrSellerConflict = errors.New("El CID ya existe")
)

// Products errors
var (
	ErrProductCodeAlreadyExists = errors.New("Codigo de producto ya existente")
	ErrProductNotFound          = errors.New("Producto no encontrado")
	ErrInvalidProductId         = errors.New("Id del producto inválido")
	ErrAllFieldsFilled          = errors.New("Todos los campos deben estar llenos")
	ErrProductAlreadyExists     = errors.New("Producto ya existe")
	ErrInvalidRequest           = errors.New("Petición inválida")
	ErrValidationError          = errors.New("Error al validar el producto")
)

// Employee errors
var (
	ErrEmployeeNotFound        = errors.New("El empleado no existe")
	ErrEmployeeAlreadyExists   = errors.New("El empleado ya existe")
	ErrCardNumberAlreadyExists = errors.New("El número de tarjeta ya existe")
)
