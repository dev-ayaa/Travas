package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/validator/v10"
	"log"
)

type Tools struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	Session     *scs.SessionManager
	Validator   *validator.Validate
}
