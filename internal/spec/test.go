package spec

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nobe4/go-cli-comparison/internal/root"
)

type Test struct {
	Args     []string
	Want     Options
	Success  bool
	location string
}

var errNoTestList = errors.New("could not find the test list")

func (t Test) Location() string {
	if t.location != "" {
		return t.location
	}

	r, err := root.Root()
	if err != nil {
		return fmt.Sprintf("could not get the root: %v", err)
	}

	path := filepath.Join(r, "internal", "spec", "list.go")

	rawContent, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("%v %s: %v", errNoTestList, path, err)
	}

	content := string(rawContent)

	args := []string{}
	for _, a := range t.Args {
		args = append(args, `"`+a+`"`)
	}

	start := 0
	end := 0
	found := false
	pattern := "Args: []string{" + strings.Join(args, ", ") + "},"

	for i, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)

		if line == "{" {
			start = i
		}

		if line == pattern {
			found = true
		}

		if found {
			if line == "}," {
				end = i

				break
			}
		}
	}

	if !found {
		return ""
	}

	t.location = fmt.Sprintf(
		"https://github.com/nobe4/go-cli-comparison/blob/main/internal/spec/list.go#L%d-L%d",
		// GitHub use 1-indexed lines
		start+1, end+1,
	)

	return t.location
}
