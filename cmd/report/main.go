package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/nobe4/go-cli-comparison/internal/format"
	"github.com/nobe4/go-cli-comparison/internal/library"
	"github.com/nobe4/go-cli-comparison/internal/result"
	"github.com/nobe4/go-cli-comparison/internal/spec"
)

const (
	pathSuffix    = "/main.go"
	root          = "libraries/"
	expectedParts = 3

	marker         = "<!-- marker:comparison-table -->"
	readmeTemplate = marker + `
| library | tests | updated | stars |
| --- | --- | --- | --- |
{{ range $i, $row := .ResultByLibs -}}
    {{- with index $.Libs $i -}}
        |[{{ .Name }}]({{ .URL }})|
        {{- range $j, $success := $row -}}
                [{{- if $success }}✅{{ else }}❌{{ end -}}]({{(index $.Tests $j).Location}})
            {{- end -}}
        |{{ .LastUpdate -}}
        |{{ .Stars }}|
    {{- end }}
{{ end }}

| test | libraries |
| --- | --- |
{{ range $i, $row := .ResultByTests -}}
    {{- with index $.Tests $i -}}
        |{{ .Args }}|
        {{- range $j, $success := $row -}}
                [{{- if $success }}✅{{ else }}❌{{ end -}}]({{(index $.Libs $j).Location}})
        {{- end -}}|
    {{- end }}
{{ end }}

` + marker
)

var (
	errNoMarker = errors.New("could not find marker comments")
	errNoToken  = errors.New("could not find the GITHUB_TOKEN environment variable")
)

type templateData struct {
	Libs          []*library.Library
	Tests         []spec.Test
	ResultByLibs  result.Result
	ResultByTests result.Result
}

type repo struct {
	PushedAt       string `json:"pushed_at"`
	StargazerCount int    `json:"stargazers_count"`
}

func main() {
	var err error

	libs, err := library.List()
	if err != nil {
		log.Fatalf("could not list libraries: %q", err)
	}

	if err := populateStats(libs); err != nil {
		panic(err)
	}

	if err := updateReadme(libs); err != nil {
		panic(err)
	}
}

func populateStats(libs []*library.Library) error {
	for _, lib := range libs {
		slog.Info("populating stats", "lib", lib.Name)

		if lib.URL == "" {
			continue
		}

		if fullName, found := strings.CutPrefix(lib.URL, "https://github.com/"); found {
			if err := populateGitHubStats(lib, fullName); err != nil {
				return fmt.Errorf("could not populate stats for %s: %q", lib.URL, err)
			}
		}
	}

	return nil
}

func populateGitHubStats(lib *library.Library, name string) error {
	slog.Info("populating stats from GitHub", "lib", lib.Name)

	req, err := http.NewRequestWithContext(
		context.TODO(),
		http.MethodGet,
		"https://api.github.com/repos/"+name,
		nil,
	)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	token, found := os.LookupEnv("GITHUB_TOKEN")
	if !found {
		return errNoToken
	}

	req.SetBasicAuth("token", token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not fetch repo '%s': %w", name, err)
	}
	defer resp.Body.Close()

	r := repo{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return fmt.Errorf("could not decode response for repo '%s': %w", name, err)
	}

	t, err := time.Parse(time.RFC3339, r.PushedAt)
	if err != nil {
		return fmt.Errorf("could not parse repo '%s' field PushedAt %q, %w", name, r.PushedAt, err)
	}

	lib.LastUpdate = format.Time(t)
	lib.Stars = format.Count(r.StargazerCount)

	return nil
}

func updateReadme(libs []*library.Library) error {
	slog.Info("updating README")

	content, err := os.ReadFile("tests/results.txt")
	if err != nil {
		return fmt.Errorf("could not read tests/results.txt: %w\ntry: go test ./...", err)
	}

	r := result.Result{}
	result.Unmarshal(content, &r)

	t, err := template.New("readme table").Parse(readmeTemplate)
	if err != nil {
		return fmt.Errorf("could not parse template: %w", err)
	}

	templateData := templateData{
		Libs:          libs,
		Tests:         spec.Tests,
		ResultByLibs:  r,
		ResultByTests: r.Rotate(),
	}

	table := bytes.Buffer{}
	if err := t.Execute(&table, templateData); err != nil {
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
