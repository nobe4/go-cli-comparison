package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/nobe4/go-cli-comparison/internal/library"
	"github.com/nobe4/go-cli-comparison/internal/relative"
)

const (
	pathSuffix    = "/main.go"
	root          = "libraries/"
	expectedParts = 3

	marker         = "<!-- marker:comparison-table -->"
	readmeTemplate = marker + `
| repo | last-commit | stars |
| --- | --- | --- |
{{ range . -}}
| [{{ .FullName }}]({{ .HTMLURL }}) | {{ .PushedAtFormatted }} | {{ .StargazerCount }} |
{{ end -}}
` + marker
)

var (
	errNoMarker = errors.New("could not find marker comments")
	errNoToken  = errors.New("could not find the GITHUB_TOKEN environment variable")
)

type repo struct {
	FullName          string `json:"full_name"`
	PushedAt          string `json:"pushed_at"`
	PushedAtFormatted string `json:"-"`
	StargazerCount    int    `json:"stargazers_count"`
	HTMLURL           string `json:"html_url"`
}

func main() {
	var err error

	var repos []repo

	if repos, err = listRepositories(); err != nil {
		panic(err)
	}

	if err := updateReadme(repos); err != nil {
		panic(err)
	}
}

func listRepositories() ([]repo, error) {
	libs, err := library.List()
	if err != nil {
		return nil, fmt.Errorf("could not list libraries: %w", err)
	}

	list := make([]repo, 0, len(libs))

	for _, lib := range libs {
		if lib.Name == "std/flag" {
			continue
		}

		r, err := fetchRepo(lib.Name)
		if err != nil {
			return nil, err
		}

		list = append(list, r)
	}

	if err != nil {
		return nil, fmt.Errorf("could not list repositories: %w", err)
	}

	return list, nil
}

func fetchRepo(name string) (repo, error) {
	req, err := http.NewRequestWithContext(
		context.TODO(),
		http.MethodGet,
		"https://api.github.com/repos/"+name,
		nil,
	)
	if err != nil {
		return repo{}, fmt.Errorf("could not create request: %w", err)
	}

	token, found := os.LookupEnv("GITHUB_TOKEN")
	if !found {
		return repo{}, errNoToken
	}

	req.SetBasicAuth("token", token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return repo{}, fmt.Errorf("could not fetch repo '%s': %w", name, err)
	}
	defer resp.Body.Close()

	r := repo{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return repo{}, fmt.Errorf("could not decode response for repo '%s': %w", name, err)
	}

	t, err := time.Parse(time.RFC3339, r.PushedAt)
	if err != nil {
		return repo{}, fmt.Errorf("could not parse repo '%s' field PushedAt %q, %w", name, r.PushedAt, err)
	}

	r.PushedAtFormatted = relative.Time(t)

	return r, nil
}

func updateReadme(repos []repo) error {
	t, err := template.New("readme table").Parse(readmeTemplate)
	if err != nil {
		return fmt.Errorf("could not parse template: %w", err)
	}

	table := bytes.Buffer{}
	if err := t.Execute(&table, repos); err != nil {
		return fmt.Errorf("could not execute template: %w", err)
	}

	readme, err := os.ReadFile("README.md")
	if err != nil {
		return fmt.Errorf("could not read README.md: %w", err)
	}

	parts := bytes.Split(readme, []byte(marker))
	if len(parts) < expectedParts {
		return errNoMarker
	}

	header, _, tail := parts[0], parts[1], parts[2]

	readme = bytes.Join([][]byte{header, table.Bytes(), tail}, []byte(""))

	if err := os.WriteFile("README.md", readme, 0o600); err != nil {
		return fmt.Errorf("could not write README.md: %w", err)
	}

	return nil
}
