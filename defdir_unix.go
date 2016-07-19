// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package config

import (
	"os"
	"path/filepath"
)

func defaultDirectory() string {
	base := filepath.Base(os.Args[0])
	ext := filepath.Ext(base)

	cnfDir := "/etc"
	name := base[0 : len(base)-len(ext)]

	return filepath.Join(cnfDir, name)
}
