# systemuri

A go library to register custom URL scheme in the operating system. Inspired by https://github.com/Vrixyz/system_uri

## About

This library allows registering a custom URL scheme like `foobar://` in the system and execute a defined program.

### Windows

The scheme is registered via a registry entry under TODO.

### Darwin

The scheme is registered via TODO.

### Linux

The scheme is registered via a scheme `(scheme)-url-handler.desktop` file under the users XDG application directory.

## Contributing ♥

See [CONTRIBUTING.md](CONTRIBUTING.md).

Made with love and go.
