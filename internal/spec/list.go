//nolint:mnd,gochecknoglobals // For testing, this is fine.
package spec

var Tests = []Test{
	{
		Want: Options{},
	},

	// A
	{
		Args: []string{"-a"},
		Want: Options{A: true},
	},

	// B
	{
		Args: []string{"-b", "-b", "-b"},
		Want: Options{B: 3},
	},

	// C
	{
		Args: []string{"-c", "a"},
		Want: Options{C: "a"},
	},
	{
		Args: []string{"-ca"},
		Want: Options{C: "a"},
	},
	{
		Args: []string{"-c=a"},
		Want: Options{C: "a"},
	},

	// D
	{
		Args: []string{"-d", "a"},
		Want: Options{D: []string{"a"}},
	},
	{
		Args: []string{"-d", "a", "-d", "a"},
		Want: Options{D: []string{"a", "a"}},
	},
	{
		Args: []string{"-d", "a", "-d", "b"},
		Want: Options{D: []string{"a", "b"}},
	},
}
