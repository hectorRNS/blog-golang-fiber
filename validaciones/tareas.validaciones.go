package validaciones

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var nombreRegex = regexp.MustCompile(`^[a-zA-Z0-9ñÑáéíóúÁÉÍÓÚ\s]+$`)

func nombreValido(fl validator.FieldLevel) bool {
	return nombreRegex.MatchString(fl.Field().String())
}

func RegistrarValiaciones(validacion *validator.Validate) {
	validacion.RegisterValidation("texto", nombreValido)
}
