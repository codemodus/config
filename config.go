// Package config provides an interface and an initialization function
// for handling configuration values stored as JSON.  The JSON structure
// is defined by a user configurable struct which implements Configurator.
// Nested configuration files can be handled so long as Init is called
// within the parent InitPost.
package config

import (
	"encoding/json"
	"os"
	"path"
)

var (
	// DefaultConfDir sets the default configuration file parent directory.
	DefaultConfDir = "/etc/" + path.Base(os.Args[0])
)

// Configurator defines the basic functionality required to work with config.
type Configurator interface {
	// ConfDir should return the absolute path to the configuration directory
	// as a string.
	ConfDir() string
	// InitPost should contain any post initialization logic to be applied to
	// the Configurator, and return any processing errors.
	InitPost() error
}

// Init decodes the provided JSON file into the provided Configurator.
func Init(c Configurator, filename string) (err error) {
	d := c.ConfDir()
	if d == "" {
		d = DefaultConfDir
	}
	if filename == "" {
		filename = "config.json"
	}

	f, err := os.Open(d + "/" + filename)
	if err != nil {
		return err
	}

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return err
	}

	err = c.InitPost()
	if err != nil {
		return err
	}
	return nil
}
