//go:build windows

package systemuri

import (
	"log"
)

func registerURLHandler(name string, scheme string, applicationPath string, argumentsPattern string) error {
	// Implement Windows registration logic
	log.Println("Registering URL handler on Windows is not implemented")
	return nil
}

func unregisterURLHandler(scheme string) error {
	log.Println("Not yet implemented")
	return nil
}

func unregisterURLHandlerByPath(applicationPath string) error {
	log.Println("Not yet implemented")
	return nil
}

//import (
//"fmt"
//"golang.org/x/sys/windows/registry"
//"os/exec"
//)

//func registerURLHandler(name string, schema string, applicationPath string) error {
//	if applicationPath == "" {
//		return fmt.Errorf("applicationPath is empty")
//	}
//
//	k, _, err := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\Classes\`+schema, registry.WRITE)
//	if err != nil {
//		return fmt.Errorf("failed to create registry key: %w", err)
//	}
//	defer k.Close()
//
//	err = k.SetStringValue("", "URL:"+name+" Protocol")
//	if err != nil {
//		return fmt.Errorf("failed to set default value: %w", err)
//	}
//
//	err = k.SetStringValue("URL Protocol", "")
//	if err != nil {
//		return fmt.Errorf("failed to set URL Protocol value: %w", err)
//	}
//
//	shellKey, _, err := registry.CreateKey(k, `shell\open\command`, registry.WRITE)
//	if err != nil {
//		return fmt.Errorf("failed to create shell\\open\\command key: %w", err)
//	}
//	defer shellKey.Close()
//
//	command := fmt.Sprintf(`"%s" "%%1"`, applicationPath)
//	err = shellKey.SetStringValue("", command)
//	if err != nil {
//		return fmt.Errorf("failed to set command value: %w", err)
//	}
//	return nil
//}
//
//func unregisterURLHandler(schema string) error {
//	err := registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\Classes\`+schema)
//	if err != nil {
//		return fmt.Errorf("failed to remove registry key: %w", err)
//	}
//	return nil
//}
