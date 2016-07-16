// Package config provides an interface and initialization function for
// handling configuration values stored as JSON.  The JSON structure is defined
// by a user configurable struct which implements Configurator.  Nest
// Configurator instances by calling Init within the parent's InitPost.
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

var (
	// DefaultDir is the configuration directory fallback.
	DefaultDir = defaultDirectory()
	// DefaultFilename is the configuration filename fallback.
	DefaultFilename = "config.cnf"
	// ErrNotUnmarshalable indicates that the relevant data is neither JSON, TOML, or YAML.
	ErrNotUnmarshalable = errors.New("provided data cannot be unmarshaled")
)

// Configurator defines the basic functionality required to work with config.
type Configurator interface {
	// InitPost should contain any post initialization logic to be applied to
	// the Configurator, and return any processing errors.
	InitPost() error
}

// Init decodes the provided JSON file into the provided Configurator.
func Init(c Configurator, file string) (err error) {
	if file == "" {
		file = path.Join(DefaultDir, DefaultFilename)
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}

	err = decode(f, c)
	if err != nil {
		return err
	}

	err = c.InitPost()

	return err
}

// Config is provided for embedding default methods needed to satisfy the
// Configurator interface.
type Config struct{}

// InitPost runs post config initialization processing.
func (c *Config) InitPost() error {
	return nil
}

func decode(f io.Reader, c Configurator) error {
	bb := &bytes.Buffer{}
	_, err := bb.ReadFrom(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bb.Bytes(), c)
	if err == nil {
		return nil
	}

	_, err = toml.Decode(bb.String(), c)
	if err == nil {
		return nil
	}

	return ErrNotUnmarshalable
}
