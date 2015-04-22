# config

    import "github.com/codemodus/config"

Package config provides an interface and an initialization function for handling
configuration values stored as JSON. The JSON structure is defined by a user
configurable struct which implements Configurator. Nested configuration files
can be handled so long as Init is called within the parent InitPost.

## Usage

View the [GoDoc](http://godoc.org/github.com/codemodus/config)
