package library

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	pathSuffix    = "/main.go"
	librariesPath = "libraries/"
)

var ErrNoProjectRoot = errors.New("could not find the project root go.mod")

type Library struct {
	Name           string
	NormalizedName string
	Path           string
}

func List() ([]Library, error) {
	root, err := ProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("could not get the project root: %w", err)
	}

	root = filepath.Join(root, librariesPath) + "/"

	var list []Library

	err = filepath.Walk(root, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, pathSuffix) {
			name := strings.TrimSuffix(path, pathSuffix)
			name = strings.TrimPrefix(name, root)

			list = append(list, Library{
				Name:           name,
				NormalizedName: strings.ReplaceAll(name, "/", "_"),
				Path:           path,
			})
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not walk from the root %s: %w", librariesPath, err)
	}

	return list, nil
}

func ProjectRoot() (string, error) {
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
