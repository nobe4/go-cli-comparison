package library

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nobe4/go-cli-comparison/internal/root"
)

const (
	pathSuffix    = "/main.go"
	librariesPath = "libraries/"
)

type Library struct {
	Name           string
	NormalizedName string
	Path           string
}

func List() ([]Library, error) {
	projectRoot, err := root.Root()
	if err != nil {
		return nil, fmt.Errorf("could not get the project root: %w", err)
	}

	projectRoot = filepath.Join(projectRoot, librariesPath) + "/"

	var list []Library

	err = filepath.Walk(projectRoot, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, pathSuffix) {
			name := strings.TrimSuffix(path, pathSuffix)
			name = strings.TrimPrefix(name, projectRoot)

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
