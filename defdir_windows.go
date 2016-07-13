// +build windows

package config

import (
	"os"
	"path"
)

func defaultDirectory() string {
	base := path.Base(os.Args[0])
	ext := path.Ext(base)

	drv := os.Getenv("SystemDrive")
	pdDir := "ProgramData"
	name := base[0 : len(base)-len(ext)]

	return path.Join(drv, pdDir, name, name)
}
