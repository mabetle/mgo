package mcore

import (
	"os"
)

// GetEnv returns env value
func GetEnv(name string) string {
	return os.Getenv(name)
}

// PutEnv puts env
func PutEnv(name, value string) error {
	return os.Setenv(name, value)
}

// IsEnvExist is env exists
func IsEnvExist(name string) bool {
	_, b := os.LookupEnv(name)
	return b
}
