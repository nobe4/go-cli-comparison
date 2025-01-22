//nolint:mnd,gochecknoglobals // For testing, this is fine.
package test

import "github.com/nobe4/go-cli-comparison/internal/spec"

var tests = []struct {
	args []string
	want spec.Options
}{
	{
		want: spec.Options{},
	},

	// A
	{
		args: []string{"-a"},
		want: spec.Options{A: true},
	},

	// B
	{
		args: []string{"-b", "-b", "-b"},
		want: spec.Options{B: 3},
	},

	// C
	{
		args: []string{"-c", "a"},
		want: spec.Options{C: "a"},
	},
	{
		args: []string{"-ca"},
		want: spec.Options{C: "a"},
	},
	{
		args: []string{"-c=a"},
		want: spec.Options{C: "a"},
	},

	// D
	{
		args: []string{"-d", "a"},
		want: spec.Options{D: []string{"a"}},
	},
	{
		args: []string{"-d", "a", "-d", "a"},
		want: spec.Options{D: []string{"a", "a"}},
	},
	{
		args: []string{"-d", "a", "-d", "b"},
		want: spec.Options{D: []string{"a", "b"}},
	},
}
