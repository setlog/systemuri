//go:build windows

package systemuri

import (
	"log"
)

func registerURLHandler(schema, applicationPath string) error {
	// Implement Windows registration logic
	log.Println("Registering URL handler on Windows is not implemented")
	return nil
}
