// Package config will manage all application level configurations
// config file will be taken based on the application environment
// all the configuration available in vault file will be overwritten
// this will be immutable as it always provides the value of the struct
package config

import (
	"db-poc/pkg/db/sql"
)

// AppConfig - global configuration struct definition
type AppConfig struct {
	Database sql.Config `toml:"database"`
}
