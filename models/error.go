package models

import (
	"errors"
)

type ValidationError error

var (
	errorUsername          = ValidationError(errors.New("El username no debe estar vacio."))
	errorShotUsername      = ValidationError(errors.New("El username es demasiado corto."))
	errorLargeUsername     = ValidationError(errors.New("El username es demasiado largo."))
	errorDuplicateUsername = ValidationError(errors.New("El username ya existe"))

	errorLogin = ValidationError(errors.New("Usuario o Password Incorrectos"))

	errorEmail              = ValidationError(errors.New("Formato invalidado"))
	errorPasswordEncryption = ValidationError(errors.New("No es posible cifrar el texto"))

	errorData  = errors.New("Mal formato de datos")
	errorDate  = errors.New("Mal formato de fecha")
	errorValue = errors.New("Mal formato de Valor")
	errorTopic = errors.New("Mal formato de Topico")

	ErrorExpiredToken    = errors.New("Token Expirado")
	ErrorValidationToken = errors.New("Error de Validacion de Token")
	ErrorSign            = errors.New("Error de Token")
)
