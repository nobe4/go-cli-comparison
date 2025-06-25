//nolint:mnd,gochecknoglobals // For testing, this is fine.
package spec

var List = Specs{
	{
		ID:          0,
		Description: "nothing",
		Default:     "",
		Regex:       "",
		Comments:    []string{"By default, nothing happens"},

		Tests: []Test{
			{
				Want: Options{},
			},
		},
	},

	{
		ID:          1,
		Description: "Boolean toggle",
		Default:     "false",
		Regex:       `-\w`,

		Tests: []Test{
			{
				Args: []string{"-a"},
				Want: Options{A: true},
			},
		},
	},

	{
		ID:          2,
		Description: "Count how many times the flag was passed",
		Default:     "0",
		Regex:       `(-\w)( \1)*`,

		Tests: []Test{
			{
				Args: []string{"-b"},
				Want: Options{B: 1},
			},
			{
				Args: []string{"-b", "-b", "-b"},
				Want: Options{B: 3},
			},
		},
	},

	{
		ID:          3,
		Description: "Store an string",
		Default:     `""`,
		Type:        "string",
		Regex:       `-\w(=| )?\.+`,

		Tests: []Test{
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
		},
	},

	{
		ID:          4,
		Description: "Store an array of strings",
		Default:     `[]`,
		Type:        "[]string",
		Regex:       `(-\w)(=| )?\.+( \1(=| )?\.+)*`,

		Tests: []Test{
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
		},
	},
}
