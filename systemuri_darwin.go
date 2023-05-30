//go:build darwin

package systemuri

import (
	"log"
)

func registerURLHandler(schema, applicationPath string) error {
	// Implement macOS registration logic
	log.Println("Registering URL handler on macOS is not implemented")
	return nil
}

func unregisterURLHandler(applicationPath string) error {
	// Implement macOS registration logic
	log.Println("Registering URL handler on macOS is not implemented")
	return nil
}
