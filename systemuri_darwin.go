//go:build darwin

package systemuri

import (
	"log"
)

func registerURLHandler(name string, scheme string, applicationPath string, argumentsPattern string) error {
	// Implement macOS registration logic
	log.Println("Registering URL handler on macOS is not implemented")
	return nil
}

func unregisterURLHandler(applicationPath string) error {
	// Implement macOS unregistration logic
	log.Println("Not yet implemented")
	return nil
}

func unregisterURLHandlerByPath(applicationPath string) error {
	log.Println("Not yet implemented")
	return nil
}
