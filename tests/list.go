package test

import "github.com/nobe4/go-cli-comparison/internal/spec"

//nolint:gochecknoglobals // This is better than a local var.
var tests = []struct {
	args []string
	want spec.Options
}{
	{
		want: spec.Options{},
	},
	{
		args: []string{"-a"},
		want: spec.Options{A: true},
	},
}
