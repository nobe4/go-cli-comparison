package library

import (
	"bytes"
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
	URL            string
	Location       string
	Path           string
	FullPath       string // full path FS root
	Stars          string
	LastUpdate     string
}

func List() ([]*Library, error) {
	projectRoot, err := root.Root()
	if err != nil {
		return nil, fmt.Errorf("could not get the project root: %w", err)
	}

	librariesPath := filepath.Join(projectRoot, librariesPath) + "/"

	var list []*Library

	err = filepath.Walk(librariesPath, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, pathSuffix) {
			file, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("could not read the file %s: %w", path, err)
			}

			name := ""
			url := ""

			for _, line := range bytes.Split(file, []byte("\n")) {
				if v, found := bytes.CutPrefix(line, []byte("Name = ")); found {
					name = string(bytes.TrimSpace(v))
				}

				if v, found := bytes.CutPrefix(line, []byte("URL  = ")); found {
					url = string(bytes.TrimSpace(v))
				}
			}

			relativePath, _ := strings.CutPrefix(path, projectRoot+"/")

			list = append(list, &Library{
				Name:           name,
				URL:            url,
				NormalizedName: strings.ReplaceAll(name, "/", "_"),
				FullPath:       path,
				Path:           relativePath,
				Location:       "https://github.com/nobe4/go-cli-comparison/blob/main/" + relativePath,
			})
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not walk from the root %s: %w", librariesPath, err)
	}

	return list, nil
}
