package config

import (
	"encoding/json"
	"os"
	"path"
)

var (
	DefaultConfDir = "/etc/" + path.Base(os.Args[0])
)

type Configurator interface {
	ConfDir() string
	InitPost() error
}

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
