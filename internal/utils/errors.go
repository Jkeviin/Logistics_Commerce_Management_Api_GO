package utils

import (
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// General error
var (
	ErrInternalServer       = errors.New("Ha ocurrido un error, contacte al administrador")
	ErrBadRequest           = errors.New("El contenido es inválido o está mal formado")
	ErrValidation           = errors.New("Ha ocurrido un error de validación")
	ErrInvalidFilter        = errors.New("Filtro inválido: se debe proporcionar al menos un campo de filtro")
	ErrMandatoryId          = errors.New("El ID es obligatorio")
	ErrInvalidId            = errors.New("El ID debe ser un número entero válido")
	ErrNoChanges            = errors.New("No se han realizado cambios en los datos almacenados")
	ErrNoDeleteByForeignKey = errors.New("No se puede eliminar el registro porque tiene registros relacionados")
	ErrFailedParseDate      = errors.New("Error al parsear la fecha de caducidad")
)

// Warehouses errors
var (
	ErrWarehouseCodeAlreadyExists = errors.New("El código del almacén ya existe")
	ErrWareHouseNotFound          = errors.New("El almacén no existe")
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
	ErrSectionForeignKey    = errors.New("La sección no existe")
)

// Sellers errors
var (
	ErrSellerNotFound = errors.New("El vendedor no existe")
	ErrSellerConflict = errors.New("El CID ya existe")
)

// Carries errors
var (
	ErrCarriesAlreadyExists = errors.New("El transportista ya existe")
	ErrCarriesNotFound      = errors.New("El transportista no existe")
	ErrCarriesForeignKey    = errors.New("El transportista no existe")
)

// Products errors
var (
	ErrProductCodeAlreadyExists = errors.New("Codigo de producto ya existente")
	ErrProductNotFound          = errors.New("Producto no existe")
	ErrInvalidProductId         = errors.New("Id del producto inválido")
	ErrProductAlreadyExists     = errors.New("Producto ya existe")
	ErrInvalidRequest           = errors.New("Petición inválida")
	ErrProductForeignKey        = errors.New("El producto no existe")
)

// ProductType errors
var (
	ErrProductTypeNotFound = errors.New("El tipo de producto no existe")
)

// Employee errors
var (
	ErrEmployeeNotFound        = errors.New("El empleado no existe")
	ErrCardNumberAlreadyExists = errors.New("El número de tarjeta ya existe")
)

// Locality errors
var (
	ErrIdAlreadyExist  = errors.New("El id (código postal) ya existe")
	ErrSellersNotFound = errors.New("No se encontraron Sellers en esa Locality")
)

// Locality errors
var (
	ErrLocalityForeignKey = errors.New("La localidad no existe")
)

// InboundOrder errors
var (
	ErrOrderNumberAlreadyExists = errors.New("El número de orden ya existe")
)

// ProductBatch errors
var (
	ErrProductBatchNotFound = errors.New("El lote de producto no existe")
)

// Purchase Orders errors
var (
	ErrPurchaseOrderConflict = errors.New("El número de orden ya existe")
)

// HandleError handles errors based on a mapping of error types to HTTP status codes
func HandleError(w http.ResponseWriter, err error, errorMapping map[error]int, defaultStatus int) {
	for errorMap, status := range errorMapping {
		if errors.Is(err, errorMap) {
			response.Error(w, status, err.Error())
			return
		}
	}
	if defaultStatus == http.StatusInternalServerError {
		response.Error(w, defaultStatus, ErrInternalServer.Error())
		return
	}
	response.Error(w, defaultStatus, err.Error())
}

// ErrorInSlice checks if an error is present in a slice of errors.
func ErrorInSlice(slice []error, err error) bool {
	for _, e := range slice {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}
