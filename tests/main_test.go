package test

import (
	"context"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nobe4/go-cli-comparison/internal/library"
	"github.com/nobe4/go-cli-comparison/internal/spec"
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
				t.Run(strings.Join(test.args, " "), func(t *testing.T) {
					got := run(t, bin, test.args)

					if !test.want.Equal(got) {
						t.Fatalf("want Options to be %v, got: %v", test.want, got)
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

func run(t *testing.T, bin string, args []string) spec.Options {
	t.Helper()

	cmd := exec.CommandContext(
		context.TODO(),
		bin,
		args...,
	)

	t.Logf("running %s with args %v", bin, args)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("out: %s", out)
		t.Fatalf("could not run: %v", err)
	}

	o, err := spec.Unmarshal(out)
	if err != nil {
		t.Fatalf("could not unmarshal output: %v", err)
	}

	return o
}
