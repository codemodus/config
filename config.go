package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

var (
	defaultDirConf = "/etc/" + path.Base(os.Args[0])
)

type Configurator interface {
	ConfDir() string
	InitPost()
}

func InitConfig(c Configurator, filename string) {
	dirConf := c.ConfDir()
	if dirConf == "" {
		dirConf = defaultDirConf
	}

	if filename == "" {
		filename = "config.json"
	}

	f, err := os.Open(dirConf + "/" + filename)
	if err != nil {
		log.Fatalln("Configuration file not found:", err)
	}

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		log.Fatalln("Configuration decoding failed:", err)
	}

	c.InitPost()
}
