package mcore

import (
	"os"
	"path/filepath"
)

// IsGitRepo reports whether the working directory is inside a Git repository.
func IsGitRepo() bool {
	p := ".git"
	for {
		fi, err := os.Stat(p)
		if os.IsNotExist(err) {
			p = filepath.Join("..", p)
			continue
		}
		if err != nil || !fi.IsDir() {
			return false
		}
		return true
	}
}
