package root

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrNoProjectRoot = errors.New("could not find the project root go.mod")

func Root() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get the working directory: %w", err)
	}

	dir = filepath.Clean(dir)

	for {
		if f, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !f.IsDir() {
			return dir, nil
		}

		newDir := filepath.Dir(dir)

		// We are at the root
		if newDir == dir {
			break
		}

		dir = newDir
	}

	return "", ErrNoProjectRoot
}
