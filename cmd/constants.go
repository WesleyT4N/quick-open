package cmd

import (
	"os"
	"path/filepath"
)

var ConfigDir = filepath.Join(os.Getenv("HOME"), ".config/quick-open")
