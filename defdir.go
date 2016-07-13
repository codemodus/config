// +build !windows

package config

import (
	"os"
	"path"
)

func defaultDirectory() string {
	base := path.Base(os.Args[0])
	ext := path.Ext(base)

	cnfDir := "/etc"
	name := base[0 : len(base)-len(ext)]

	return path.Join(cnfDir, name)
}
