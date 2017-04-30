// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package config

import (
	"os"
	"path/filepath"
)

func defaultDirectory(dirRoot string) string {
	base := filepath.Base(os.Args[0])
	ext := filepath.Ext(base)

	name := base[0 : len(base)-len(ext)]

	return filepath.Join("/", dirRoot, name)
}

func defaultConfDir() string {
	return defaultDirectory("/etc")
}

func defaultLibDir() string {
	return defaultDirectory("/var/lib")
}
