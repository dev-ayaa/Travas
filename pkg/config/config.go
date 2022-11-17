package config

import (
	"github.com/alexedwards/scs/v2"
	"log"
)

type TravasConfig struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	Session     *scs.SessionManager
}
