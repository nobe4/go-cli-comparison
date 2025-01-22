package test

import (
	"context"
	"io"
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
				t.Run(strings.Join(test.Args, " "), func(t *testing.T) {
					got := run(t, bin, test.Args)

					t.Log(test.Location())

					if !test.Want.Equal(got) {
						t.Fatalf("want Options to be %v, got: %v", test.Want, got)
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

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		t.Fatalf("could not get stderr pipe: %v", err)
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("could not get stdout pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		t.Fatalf("could not run: %v", err)
	}

	stdout, err := io.ReadAll(stdoutPipe)
	if err != nil {
		t.Logf("could not read stdout: %v", err)
	}

	stderr, err := io.ReadAll(stderrPipe)
	if err != nil {
		t.Logf("could not read stderr: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		t.Logf("stderr:\n%s", stderr)
		t.Fatalf("could not wait for: %v", err)
	}

	t.Logf("stderr:\n%s", stderr)
	t.Logf("stdout:\n%s", stdout)

	o, err := spec.Unmarshal(stdout)
	if err != nil {
		t.Fatalf("could not unmarshal output '%s': %v", stdout, err)
	}

	return o
}
