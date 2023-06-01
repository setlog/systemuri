package systemuri

// RegisterURLHandler registers a custom URL handler for the specified schema and application. argumentsPattern the %s is replaced by the system-specific placeholder
func RegisterURLHandler(name string, scheme string, applicationPath string, argumentsPattern string) error {
	return registerURLHandler(name, scheme, applicationPath, argumentsPattern)
}

// UnregisterURLHandler removes all registered entries based on the scheme
func UnregisterURLHandler(scheme string) error {
	return unregisterURLHandler(scheme)
}

// UnregisterURLHandlerByPath removes all registered entries based on the application path
func UnregisterURLHandlerByPath(applicationPath string) error {
	return unregisterURLHandlerByPath(applicationPath)
}
