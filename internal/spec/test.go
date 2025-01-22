package spec

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nobe4/go-cli-comparison/internal/library"
)

type Test struct {
	Args []string
	Want Options
}

var errNoTestList = errors.New("could not find the test list")

// Link returns the lines containing the test.
func (t Test) Location() (string, error) {
	root, err := library.ProjectRoot()
	if err != nil {
		return "", library.ErrNoProjectRoot
	}

	path := filepath.Join(root, "tests", "list.go")

	rawContent, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("%w %s: %w", errNoTestList, path, err)
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
		return "", nil
	}

	return fmt.Sprintf("https://github.com/nobe4/go-cli-comparison/blob/main/tests/list.go#L%d-L%d", start, end), nil
}
