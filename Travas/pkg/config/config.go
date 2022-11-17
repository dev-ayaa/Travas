package config

import "log"

type TravasConfig struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}