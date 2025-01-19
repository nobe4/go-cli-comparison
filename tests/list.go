package test

//nolint:gochecknoglobals // This is better than a local var.
var tests = []struct {
	id   string
	args []string
	want string
}{
	{
		id:   "0",
		want: "aFlag: 0\n",
	},
	{
		id:   "1",
		args: []string{"-a", "1"},
		want: "aFlag: 1\n",
	},
}
