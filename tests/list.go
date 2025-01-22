//nolint:mnd,gochecknoglobals // For testing, this is fine.
package test

import "github.com/nobe4/go-cli-comparison/internal/spec"

var tests = []spec.Test{
	{
		Want: spec.Options{},
	},

	// A
	{
		Args: []string{"-a"},
		Want: spec.Options{A: true},
	},

	// B
	{
		Args: []string{"-b", "-b", "-b"},
		Want: spec.Options{B: 3},
	},

	// C
	{
		Args: []string{"-c", "a"},
		Want: spec.Options{C: "a"},
	},
	{
		Args: []string{"-ca"},
		Want: spec.Options{C: "a"},
	},
	{
		Args: []string{"-c=a"},
		Want: spec.Options{C: "a"},
	},

	// D
	{
		Args: []string{"-d", "a"},
		Want: spec.Options{D: []string{"a"}},
	},
	{
		Args: []string{"-d", "a", "-d", "a"},
		Want: spec.Options{D: []string{"a", "a"}},
	},
	{
		Args: []string{"-d", "a", "-d", "b"},
		Want: spec.Options{D: []string{"a", "b"}},
	},
}
