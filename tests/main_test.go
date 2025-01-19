package test

import (
	"context"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/nobe4/go-cli-comparison/internal/library"
)

func TestMain(t *testing.T) {
	t.Parallel()

	libs, err := library.List()
	if err != nil {
		t.Fatalf("could not get the list of libraries: %v", err)
	}

	for _, lib := range libs {
		t.Run(lib.Name, func(t *testing.T) {
			t.Parallel()

			bin := build(t, lib)

			for _, test := range tests {
				t.Run(test.id, func(t *testing.T) {
					got := run(t, bin, test.args)

					if got != test.want {
						t.Fatalf("expected output to be '%s', got: '%s'", test.want, got)
					}
				})
			}
		})
	}
}

func build(t *testing.T, lib library.Library) string {
	t.Helper()

	tmp := t.TempDir()
	bin := filepath.Join(tmp, lib.NormalizedName)

	// build executable
	cmd := exec.CommandContext(
		context.TODO(),
		"go",
		"build",
		"-o",
		bin,
		lib.Path,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("out: %s", out)
		t.Fatalf("could not build %s: %v", lib.Path, err)
	}

	return bin
}

func run(t *testing.T, bin string, args []string) string {
	t.Helper()

	cmd := exec.CommandContext(
		context.TODO(),
		bin,
		args...,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("out: %s", out)
		t.Fatalf("could not run %s: %v", bin, err)
	}

	return string(out)
}
