# systemuri

A go library to register custom URL schema in the operating system. Inspired by https://github.com/Vrixyz/system_uri

## About

This library allows registering a custom URL schema like `foobar://` in the system and execute a defined program.

### Windows

The schema is registered via a registry entry under TODO.

### Darwin

The schema is registered via TODO.

### Linux

The schema is registered via a schema `(schema)-url-handler.desktop` file under the users XDG application directory.

## Contributing ♥

See [CONTRIBUTING.md](CONTRIBUTING.md).

Made with love and go.
