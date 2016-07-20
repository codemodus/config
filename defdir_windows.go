package config

import (
	"os"
	"path/filepath"
)

func defaultDirectory() string {
	base := filepath.Base(os.Args[0])
	ext := filepath.Ext(base)

	drv := os.Getenv("SystemDrive")
	pdDir := "ProgramData"
	name := base[0 : len(base)-len(ext)]

	return filepath.Join(drv, pdDir, name, name)
}
