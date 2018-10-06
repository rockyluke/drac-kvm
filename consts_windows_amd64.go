// -*- go -*-

package main

import (
	"os"
	"path/filepath"
)

var (
	javaRoots   = []string{"C:\\Program Files\\Java\\", "C:\\Program Files (x86)\\Java\\"}
	winJavaPath = "" //C:\\Program Files (x86)\\Java\\jre7\\bin\\javaws.exe
)

// DefaultJavaPath is the default Java path on Windows
func DefaultJavaPath() string {
	//javaws lives in a folder with a version number on Windows, so we can search for javaws.exe
	for _, r := range javaRoots {
		if _, err := os.Stat(r); os.IsNotExist(err) {
			continue
		}
		filepath.Walk(r, visit)
	}
	return winJavaPath
}

func visit(path string, f os.FileInfo, err error) error {
	if winJavaPath == "" && f.Name() == "javaws.exe" {
		winJavaPath = path
	}
	return nil
}

// EOF
