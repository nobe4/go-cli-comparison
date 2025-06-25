package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"text/template"

	"github.com/nobe4/go-cli-comparison/internal/library"
	"github.com/nobe4/go-cli-comparison/internal/result"
	"github.com/nobe4/go-cli-comparison/internal/spec"
)

const (
	specsMarker   = "<!-- marker:specs-table -->"
	specsTemplate = specsMarker + `
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
        |[{{- range .Args }}{{ . }} {{ end -}}]({{ .Location }})|
        {{- range $j, $success := $row -}}
                [{{- if $success }}✅{{ else }}❌{{ end -}}]({{(index $.Libs $j).Location}})
        {{- end -}}|
    {{- end }}
{{ end }}

` + specsMarker
)

func updateSpecs() error {
	var err error

	tables := bytes.Buffer{}
	if err := t.Execute(&tables, templateData); err != nil {
		return fmt.Errorf("could not execute template: %w", err)
	}

	readme, err := os.ReadFile("README.md")
	if err != nil {
		return fmt.Errorf("could not read README.md: %w", err)
	}

	parts := bytes.Split(readme, []byte(readmeMarker))
	if len(parts) < expectedParts {
		return errNoMarker
	}

	header, _, tail := parts[0], parts[1], parts[2]

	readme = bytes.Join([][]byte{header, tables.Bytes(), tail}, []byte(""))

	if err := os.WriteFile("README.md", readme, 0o600); err != nil {
		return fmt.Errorf("could not write README.md: %w", err)
	}

	return nil
}
