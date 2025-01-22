package test

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nobe4/go-cli-comparison/internal/library"
	"github.com/nobe4/go-cli-comparison/internal/result"
	"github.com/nobe4/go-cli-comparison/internal/spec"
)

func TestMain(t *testing.T) {
	t.Parallel()

	libs, err := library.List()
	if err != nil {
		t.Fatalf("could not get the list of libraries: %v", err)
	}

	results := result.New(len(libs), len(tests))

	for i, lib := range libs {
		t.Run(lib.Name, func(t *testing.T) {
			t.Parallel()

			bin := build(t, lib)

			for j, test := range tests {
				t.Run(strings.Join(test.Args, " "), func(t *testing.T) {
					got, success := run(t, bin, test.Args)

					t.Log(test.Location())

					if success {
						success = test.Want.Equal(got)
					}

					results[i][j] = success

					log.Printf("success: %v", lib)

					if !success {
						t.Fatalf("want Options to be %v, got: %v", test.Want, got)
					}
				})
			}
		})
	}

	t.Cleanup(func() {
		if err := os.WriteFile("results.txt", results.Marshal(), 0o600); err != nil {
			t.Fatalf("could not write results to file: %v", err)
		}
	})
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

func run(t *testing.T, bin string, args []string) (spec.Options, bool) {
	t.Helper()

	cmd := exec.CommandContext(
		context.TODO(),
		bin,
		args...,
	)

	t.Logf("running %s with args %v", bin, args)

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		t.Logf("could not get stderr pipe: %v", err)

		return spec.Options{}, false
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Logf("could not get stdout pipe: %v", err)

		return spec.Options{}, false
	}

	if err := cmd.Start(); err != nil {
		t.Logf("could not run: %v", err)

		return spec.Options{}, false
	}

	stdout, err := io.ReadAll(stdoutPipe)
	if err != nil {
		t.Logf("could not read stdout: %v", err)

		return spec.Options{}, false
	}

	stderr, err := io.ReadAll(stderrPipe)
	if err != nil {
		t.Logf("could not read stderr: %v", err)

		return spec.Options{}, false
	}

	if err := cmd.Wait(); err != nil {
		t.Logf("stderr:\n%s", stderr)
		t.Logf("could not wait for: %v", err)

		return spec.Options{}, false
	}

	t.Logf("stderr:\n%s", stderr)
	t.Logf("stdout:\n%s", stdout)

	o, err := spec.Unmarshal(stdout)
	if err != nil {
		t.Logf("could not unmarshal output '%s': %v", stdout, err)

		return spec.Options{}, false
	}

	return o, true
}
