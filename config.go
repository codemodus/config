// Package config provides an interface and init function for handling
// configuration values stored as JSON.  The JSON structure is defined
// by a user configurable struct which implements Configurator.  Nested
// configuration files can be handled so long as InitConfig() is called within
// the parent InitPost().
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
	ConfDir() string
	InitPost() error
}

// InitConfig decodes the provided JSON file into the provided Configurator.
func InitConfig(c Configurator, filename string) (err error) {
	dirConf := c.ConfDir()
	if dirConf == "" {
		dirConf = DefaultConfDir
	}

	if filename == "" {
		filename = "config.json"
	}

	f, err := os.Open(dirConf + "/" + filename)
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
