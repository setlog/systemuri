package systemuri

// RegisterURLHandler registers a custom URL handler for the specified schema and application. argumentsPattern the %s is replaced by the system-specific placeholder
func RegisterURLHandler(name string, scheme string, applicationPath string, argumentsPattern string) error {
	return registerURLHandler(name, scheme, applicationPath, argumentsPattern)
}

func UnregisterURLHandler(scheme string) error {
	return unregisterURLHandler(scheme)
}
